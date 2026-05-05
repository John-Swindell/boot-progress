# 07 — Goroutines & Channels

> **Goal:** start using Go's killer feature. Spawn goroutines, talk over channels,
> coordinate with `select`. This is why Go exists.

## What you'll learn

- `go fn()` — spawning a goroutine
- Channels: `make(chan T)`, send `ch <- x`, receive `x := <-ch`
- Buffered vs unbuffered channels (and what each communicates)
- Channel direction in signatures: `chan<- T` (send-only), `<-chan T` (recv-only)
- Closing channels — who closes, why, and the comma-ok recv idiom
- `select` for multi-channel waits
- Patterns: pipeline, fan-out, fan-in, timeout, done channel

## Coming from Python

| Python | Go | Note |
|---|---|---|
| `threading.Thread(target=f).start()` | `go f()` | Cheap (~2KB stack), real parallelism (no GIL). |
| `queue.Queue` | `chan T` | Native, typed, can be buffered or sync. |
| `q.put(x)` | `ch <- x` | Sends. |
| `q.get()` | `x := <-ch` | Receives. |
| `q.put(None); ...; if x is None: break` | `close(ch)`; `for x := range ch { }` | The closed-channel idiom is the Go equivalent of "poison-pill". |
| `concurrent.futures.as_completed` | `select { case ... }` | Multiplex N channels. |

## The two channel rules

1. **The sender closes** the channel. Receivers detect "no more data" via the
   range loop ending or `v, ok := <-ch` returning `ok == false`.
2. **Don't send on a closed channel** — it panics. Don't close twice — it panics.

If you have multiple senders, none of them should close the channel; have a
coordinator close it after all senders are done (use `sync.WaitGroup` — chapter 08).

## Unbuffered vs buffered

```go
sync   := make(chan int)        // unbuffered: send blocks until someone receives
async  := make(chan int, 10)    // buffered cap 10: send blocks only when full
```

**Unbuffered = synchronization.** A successful send means a receive happened.
**Buffered = decoupling.** Sender doesn't block until the buffer fills.

Default to unbuffered. Reach for buffered when you have a real reason (e.g. fixed
worker pool, or you want to drop work when behind).

## select

```go
select {
case x := <-ch1:
    // received from ch1
case ch2 <- y:
    // sent on ch2
case <-time.After(time.Second):
    // 1 second timeout
case <-ctx.Done():
    // cancellation
}
```

Picks whichever case is ready first. With `default:`, becomes non-blocking.

## Pipeline pattern

```go
gen := func(xs ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, x := range xs { out <- x }
    }()
    return out
}

square := func(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for x := range in { out <- x*x }
    }()
    return out
}

for v := range square(gen(1, 2, 3)) { fmt.Println(v) }   // 1 4 9
```

## Your turn

```sh
go test ./07-goroutines-channels -race
```

The `-race` flag turns on the data race detector — use it religiously when
working with concurrent code. It catches bugs in dev that would silently
corrupt prod.
