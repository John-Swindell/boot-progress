# Go Lab — Zero to Hero (Platform / SRE / Kubernetes)

A self-paced, test-driven Go curriculum aimed at platform engineering and SRE work.
You write the code, the tests grade you, the scoreboard tracks your progress.

You already know Python. Every chapter starts with a "Coming from Python" callout
so you can map what you know to Go's worldview.

## How the lab works

1. **Read** the chapter `README.md`.
2. **Open** `exercises.go` and fill in the `// TODO` bodies.
3. **Run** `go test ./NN-chapter-name` until everything is green.
4. **Peek** at `solutions/exercises.go` *only after you've tried* — it's the reference, not the spec.
5. **Run** `./scripts/progress.sh` to see the scoreboard.

Whole-lab smoke check from the repo root:

```sh
go test ./...
```

## Roadmap

| # | Chapter | Concept |
|---|---------|---------|
| 00 | [setup](./00-setup) | `go mod`, `go run`, `go build`, `go test`, `go fmt`, `go vet` |
| 01 | [syntax-basics](./01-syntax-basics) | vars, types, zero values, control flow, the only-loop |
| 02 | [functions-errors](./02-functions-errors) | multi-return, `defer`, `error`, `errors.Is/As`, `panic`/`recover` |
| 03 | [collections](./03-collections) | arrays vs slices (and the aliasing trap), maps, `range` |
| 04 | [structs-methods](./04-structs-methods) | value vs pointer receivers, embedding (composition over inheritance) |
| 05 | [interfaces](./05-interfaces) | implicit satisfaction, `io.Reader`/`io.Writer`, type switches |
| 06 | [pointers-memory](./06-pointers-memory) | `&` / `*`, escape analysis intuition, when to pass a pointer |
| 07 | [goroutines-channels](./07-goroutines-channels) | `go`, buffered vs unbuffered, `select`, fan-in / fan-out |
| 08 | [sync-context](./08-sync-context) | `sync.Mutex`, `WaitGroup`, `Once`, `context.Context`, `-race` |
| 09 | [stdlib-tour](./09-stdlib-tour) | `fmt`, `strings`, `strconv`, `time`, `encoding/json` |
| 10 | [files-cli](./10-files-cli) | `os`, `io`, `bufio`, `flag`, `os/exec` — write a `wc`-clone |
| 11 | [http](./11-http) | `net/http` client + server, JSON APIs, graceful shutdown |
| 12 | [testing](./12-testing) | table tests, subtests, benchmarks, fuzzing, `httptest` |
| 13 | [modules-tooling](./13-modules-tooling) | `go mod tidy`, `replace`, `ldflags` versioning, `golangci-lint` |
| 14 | [platform-essentials](./14-platform-essentials) | `log/slog`, `os/signal` (SIGTERM), env config, retries+backoff |
| 15 | [observability](./15-observability) | Prometheus metrics, `/metrics`, `net/http/pprof`, OTel intro |
| 16 | [kubernetes-go](./16-kubernetes-go) | `client-go`: kubeconfig, list pods, stream logs, watch events |
| 17 | [**capstone: kubetail**](./17-capstone) | multi-pod log tailer — your portfolio piece |

## Coming from Python — top 10 things that will bite you

| Python | Go | Note |
|---|---|---|
| `if x:` (truthy) | `if x != "" / != 0 / != nil` | Go has no truthiness. Be explicit. |
| `None` | `nil` | But `nil` has a *type* — `var s []int` is a non-nil-but-empty slice scenario; check carefully. |
| `try / except` | `if err != nil { return err }` | Errors are values. You'll write this 1000 times. Embrace it. |
| Duck typing | Implicit interfaces | Types satisfy interfaces *implicitly* — no `implements` keyword. |
| `list` | `[]T` slice | Slices alias their backing array. `s2 := s1[2:5]` writes to `s2` mutate `s1`. |
| `dict` | `map[K]V` | Iteration order is **randomized** — by design. Don't rely on order. |
| Class inheritance | Struct embedding | Composition, not inheritance. There is no `extends`. |
| `for x in xs:` | `for i, x := range xs` | `range` returns *index, value*. Forget the index with `_, x`. |
| `__init__` | Just construct: `&Thing{Field: v}` | No constructors. Sometimes a `NewThing()` factory by convention. |
| Threads (GIL) | Goroutines | Real parallelism. Use `-race` flag religiously. |

## Setup

```sh
go version           # should be 1.22+
go mod download      # pull deps
go test ./...        # smoke check
./scripts/progress.sh
```

## Capstone: `kubetail`

When all 17 chapters are green, you build [`kubetail`](./17-capstone) — a
production-shaped CLI that streams logs from every pod matching a label
selector across a namespace, in color, multiplexed, with auto-attach for
pods born after you started watching. It exercises every concept in the lab.
