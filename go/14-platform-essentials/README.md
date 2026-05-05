# 14 — Platform Essentials

> **Goal:** the four things every long-running platform process does:
> structured logs, signal-based shutdown, config from env, retry with backoff.

## What you'll learn

- `log/slog` — structured logging built into the stdlib (since Go 1.21)
- `os/signal.NotifyContext` — turning SIGTERM/SIGINT into a cancelable context
- Twelve-factor config — read env vars at startup, fail loudly on bad values
- Retry with exponential backoff that respects context

## Coming from Python

| Python | Go | Note |
|---|---|---|
| `logging.getLogger(__name__)` | `slog.Default()` or build your own `*slog.Logger` | slog handles JSON output natively. |
| `logging.info("msg", extra={...})` | `slog.Info("msg", "key", value, ...)` | Key/value pairs as variadic args. |
| `signal.signal(SIGTERM, handler)` | `signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)` | Returns a derived ctx that cancels when a signal arrives. |
| `os.environ["X"]` (KeyError) | `os.Getenv("X")` returns "" | No exceptions; check explicitly. |
| `dotenv` | Just read `os.Getenv` — or use `github.com/caarlos0/env` for tags | Stdlib only is the default. |
| `tenacity.retry` | hand-rolled loop with `time.Sleep` (or `cenkalti/backoff`) | Stdlib retry is a 10-line function. |

## slog primer

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelInfo,
}))
logger.Info("server started", "addr", ":8080", "pid", os.Getpid())
// {"time":"2024-...","level":"INFO","msg":"server started","addr":":8080","pid":12345}
```

Level helpers: `Debug`, `Info`, `Warn`, `Error`. Attach context with `logger.With("k","v")`
to derive a child logger. Replace `slog.SetDefault(logger)` and now `slog.Info(...)`
uses your handler globally.

JSON logs are non-negotiable in containerized environments — every log
aggregator (Loki, Elasticsearch, Datadog, CloudWatch) parses them natively.

## Signal-driven shutdown

```go
func main() {
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

    if err := run(ctx); err != nil {
        slog.Error("exit", "err", err)
        os.Exit(1)
    }
}

func run(ctx context.Context) error {
    // when SIGTERM arrives, ctx.Done() closes; honor it everywhere.
    <-ctx.Done()
    return nil
}
```

This pattern is the spine of every daemon you'll write — including the capstone.

## Env config

```go
addr := os.Getenv("ADDR")
if addr == "" { addr = ":8080" }

n, err := strconv.Atoi(os.Getenv("WORKERS"))
if err != nil || n <= 0 { n = 4 }
```

Be loud about misconfig: log and exit non-zero rather than silently using a
default for required vars.

## Retry with backoff

A real retry loop: caps attempts, exponential delay, respects cancellation,
returns the last error wrapped:

```go
for attempt := 0; attempt < attempts; attempt++ {
    err := fn(ctx)
    if err == nil { return nil }
    if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
        return err  // don't retry on cancel
    }
    select {
    case <-ctx.Done():
        return ctx.Err()
    case <-time.After(backoff(attempt)):
    }
}
return fmt.Errorf("retry exhausted: %w", lastErr)
```

## Your turn

```sh
go test ./14-platform-essentials
```
