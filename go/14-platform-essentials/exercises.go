package platform

import (
	"context"
	"io"
	"log/slog"
	"time"
)

// ----- JSON Logger -----

// NewJSONLogger returns a *slog.Logger writing JSON to w at level lvl.
//
// Hint: slog.NewJSONHandler with &slog.HandlerOptions{Level: lvl}.
func NewJSONLogger(w io.Writer, lvl slog.Level) *slog.Logger {
	// TODO
	return nil
}

// ----- Env config -----

// Config carries process configuration loaded from env vars.
type Config struct {
	Addr    string        // ADDR, default ":8080"
	Workers int           // WORKERS, default 4. Must be > 0; on bad input, return error.
	Timeout time.Duration // TIMEOUT, default 5s. Use time.ParseDuration. Bad input -> error.
	Debug   bool          // DEBUG, true if env value is one of "1", "true", "TRUE", "yes". Anything else -> false. Empty -> default false.
}

// LoadConfig reads env via the supplied lookup func (so it's testable without
// touching the real environment). Pass os.Getenv when wiring this up in main.
func LoadConfig(getenv func(string) string) (Config, error) {
	// TODO
	return Config{}, nil
}

// ----- Retry with backoff (context-aware) -----

// RetryWithBackoff runs fn(ctx) until it returns nil OR attempts have been
// exhausted OR ctx is canceled.
//
// Between attempts, sleep for ExpoBackoff(attempt-1, base, max). The sleep
// itself must respect ctx (use a select over time.NewTimer(d).C and ctx.Done).
//
// On exhaustion, return the LAST error wrapped with "retry exhausted: %w".
// On ctx cancel mid-sleep, return ctx.Err() (do not wrap it).
//
// You may copy ExpoBackoff from chapter 09 — feel free to import it; here
// re-implement inline for practice.
func RetryWithBackoff(
	ctx context.Context,
	attempts int,
	base, max time.Duration,
	fn func(context.Context) error,
) error {
	// TODO
	return nil
}
