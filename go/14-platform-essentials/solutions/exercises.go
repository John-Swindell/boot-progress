package solutions

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"time"
)

func NewJSONLogger(w io.Writer, lvl slog.Level) *slog.Logger {
	return slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{Level: lvl}))
}

type Config struct {
	Addr    string
	Workers int
	Timeout time.Duration
	Debug   bool
}

func LoadConfig(getenv func(string) string) (Config, error) {
	cfg := Config{
		Addr:    ":8080",
		Workers: 4,
		Timeout: 5 * time.Second,
	}
	if v := getenv("ADDR"); v != "" {
		cfg.Addr = v
	}
	if v := getenv("WORKERS"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n <= 0 {
			return Config{}, fmt.Errorf("invalid WORKERS %q", v)
		}
		cfg.Workers = n
	}
	if v := getenv("TIMEOUT"); v != "" {
		d, err := time.ParseDuration(v)
		if err != nil {
			return Config{}, fmt.Errorf("invalid TIMEOUT %q: %w", v, err)
		}
		cfg.Timeout = d
	}
	switch getenv("DEBUG") {
	case "1", "true", "TRUE", "yes":
		cfg.Debug = true
	}
	return cfg, nil
}

func expoBackoff(attempt int, base, max time.Duration) time.Duration {
	if attempt < 0 || base <= 0 || max <= 0 {
		return 0
	}
	d := base << attempt
	if d <= 0 || d > max {
		return max
	}
	return d
}

func RetryWithBackoff(
	ctx context.Context,
	attempts int,
	base, max time.Duration,
	fn func(context.Context) error,
) error {
	if attempts <= 0 {
		return errors.New("retry: attempts must be positive")
	}
	var last error
	for i := 0; i < attempts; i++ {
		if i > 0 {
			t := time.NewTimer(expoBackoff(i-1, base, max))
			select {
			case <-ctx.Done():
				t.Stop()
				return ctx.Err()
			case <-t.C:
			}
		}
		if err := fn(ctx); err == nil {
			return nil
		} else {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return err
			}
			last = err
		}
	}
	return fmt.Errorf("retry exhausted: %w", last)
}
