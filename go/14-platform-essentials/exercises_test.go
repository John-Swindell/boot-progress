package platform

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"strings"
	"testing"
	"time"
)

func TestNewJSONLogger(t *testing.T) {
	var buf bytes.Buffer
	log := NewJSONLogger(&buf, slog.LevelInfo)
	log.Info("hi", "k", "v", "n", 42)

	var got map[string]any
	if err := json.Unmarshal(buf.Bytes(), &got); err != nil {
		t.Fatalf("not JSON: %v\n%s", err, buf.String())
	}
	if got["msg"] != "hi" || got["k"] != "v" {
		t.Errorf("got %v", got)
	}

	// Below-level should be filtered.
	buf.Reset()
	log.Debug("nope")
	if buf.Len() != 0 {
		t.Errorf("debug leaked: %q", buf.String())
	}
}

func TestLoadConfig(t *testing.T) {
	t.Run("defaults", func(t *testing.T) {
		cfg, err := LoadConfig(func(string) string { return "" })
		if err != nil {
			t.Fatalf("err = %v", err)
		}
		if cfg.Addr != ":8080" || cfg.Workers != 4 || cfg.Timeout != 5*time.Second || cfg.Debug {
			t.Errorf("got %+v", cfg)
		}
	})

	t.Run("override all", func(t *testing.T) {
		env := map[string]string{
			"ADDR":    ":9090",
			"WORKERS": "16",
			"TIMEOUT": "30s",
			"DEBUG":   "true",
		}
		cfg, err := LoadConfig(func(k string) string { return env[k] })
		if err != nil {
			t.Fatalf("err = %v", err)
		}
		if cfg.Addr != ":9090" || cfg.Workers != 16 || cfg.Timeout != 30*time.Second || !cfg.Debug {
			t.Errorf("got %+v", cfg)
		}
	})

	t.Run("bad workers", func(t *testing.T) {
		env := map[string]string{"WORKERS": "lots"}
		if _, err := LoadConfig(func(k string) string { return env[k] }); err == nil {
			t.Error("expected error")
		}
	})

	t.Run("bad timeout", func(t *testing.T) {
		env := map[string]string{"TIMEOUT": "ages"}
		if _, err := LoadConfig(func(k string) string { return env[k] }); err == nil {
			t.Error("expected error")
		}
	})
}

func TestRetryWithBackoff(t *testing.T) {
	t.Run("succeeds eventually", func(t *testing.T) {
		calls := 0
		err := RetryWithBackoff(context.Background(), 5,
			1*time.Millisecond, 10*time.Millisecond,
			func(ctx context.Context) error {
				calls++
				if calls < 3 {
					return errors.New("flake")
				}
				return nil
			})
		if err != nil {
			t.Errorf("err = %v", err)
		}
		if calls != 3 {
			t.Errorf("calls = %d, want 3", calls)
		}
	})

	t.Run("exhausts and wraps", func(t *testing.T) {
		boom := errors.New("boom")
		err := RetryWithBackoff(context.Background(), 3,
			1*time.Millisecond, 5*time.Millisecond,
			func(ctx context.Context) error { return boom })
		if !errors.Is(err, boom) {
			t.Errorf("err = %v, want errors.Is boom", err)
		}
		if !strings.Contains(err.Error(), "retry exhausted") {
			t.Errorf("missing prefix: %v", err)
		}
	})

	t.Run("respects context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(20 * time.Millisecond)
			cancel()
		}()
		err := RetryWithBackoff(ctx, 100,
			50*time.Millisecond, 200*time.Millisecond,
			func(ctx context.Context) error { return errors.New("flake") })
		if !errors.Is(err, context.Canceled) {
			t.Errorf("err = %v, want context.Canceled", err)
		}
	})
}
