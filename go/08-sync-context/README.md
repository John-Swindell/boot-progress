# 08 — sync, context, and the race detector

> **Goal:** safely share state and propagate cancellation. Every server, worker,
> and CLI you'll write in chapters 11+ depends on this.

## What you'll learn

- `sync.Mutex` / `sync.RWMutex`
- `sync.WaitGroup` for "wait for N goroutines"
- `sync.Once` for "do this exactly once"
- `context.Context` — cancellation, deadlines, request-scoped values
- The `-race` flag, and why CI without it is negligence

## Coming from Python

| Python | Go | Note |
|---|---|---|
| `threading.Lock` | `sync.Mutex` | Same idea. `Lock()`, `defer Unlock()`. |
| `threading.Event` | `chan struct{}` (closed = signal) or `context.Context` | Idiom: `<-ctx.Done()`. |
| `concurrent.futures.wait(...)` | `sync.WaitGroup` + `wg.Wait()` | Or: close the result channel after a fan-in. |
| `functools.lru_cache(maxsize=1)` | `sync.Once` + closure-captured result | Run-once initialization is a one-line `sync.Once`. |
| GIL | None | Real parallelism. **Real race conditions.** Run with `-race`. |

## context.Context — the SRE Swiss army knife

Every long-running operation, every HTTP handler, every kubernetes call takes a
`context.Context` as its **first** argument. It carries:

- **Cancellation**: caller calls `cancel()` → `ctx.Done()` channel closes → callees
  unwind. The mechanism for graceful shutdown.
- **Deadline / timeout**: `context.WithTimeout(ctx, 5*time.Second)` returns a
  child context that cancels itself when 5s elapse.
- **Values**: request-scoped data (request ID, auth subject). Use sparingly —
  it's an escape hatch, not a parameter-passing mechanism.

```go
ctx, cancel := context.WithTimeout(parent, 2*time.Second)
defer cancel()                                      // always defer cancel

req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
resp, err := http.DefaultClient.Do(req)             // honors ctx automatically
```

The "respect the context" idiom inside a loop:

```go
for {
    select {
    case <-ctx.Done():
        return ctx.Err()         // returns context.Canceled or DeadlineExceeded
    case work := <-jobs:
        // do work
    }
}
```

## sync.Mutex pattern

```go
type SafeMap struct {
    mu sync.Mutex
    m  map[string]int
}

func (s *SafeMap) Set(k string, v int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.m[k] = v
}
```

Always `defer Unlock` immediately after `Lock` — cleanest, even on panic.

`sync.RWMutex` lets many readers OR one writer hold the lock. Use it when
reads vastly outnumber writes (most caches).

## sync.WaitGroup pattern

```go
var wg sync.WaitGroup
for _, x := range items {
    wg.Add(1)                     // BEFORE go fn(); otherwise races
    go func(x int) {
        defer wg.Done()
        process(x)
    }(x)
}
wg.Wait()
```

Three rules: `Add` before `go`, `Done` deferred, `Wait` after.

## The race detector

```sh
go test -race ./08-sync-context
go run -race ./11-http
```

Adds runtime instrumentation that detects unsynchronized concurrent accesses.
Slower (~10x) but invaluable. **Run all tests with `-race` in CI.**

## Your turn

```sh
go test -race ./08-sync-context
```
