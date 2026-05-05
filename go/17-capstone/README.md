# 17 — Capstone: `kubetail`

> A real, useful, single-binary CLI: tail logs from every pod matching a label
> selector across a namespace, color-coded and multiplexed, with auto-attach
> for new pods.

You've cleared every chapter. Now you write the thing.

This capstone exercises the entire curriculum:

- `flag` (chapter 10), `slog` and signal-driven shutdown (chapter 14)
- goroutines, channels, fan-in, `select` (chapters 7–8)
- `context.Context` cancellation propagation (chapter 8)
- `client-go` list + watch + log streaming (chapter 16)
- table-driven tests with a fake clientset (chapter 12 + 16)
- `Makefile`, build flags, `golangci-lint` (chapter 13)

## What it does

```
kubetail -n <namespace> [-l <label-selector>] [--container <name>]
         [--since 5m] [--grep <regex>] [--json] [--no-color] [--kubeconfig <path>]
```

- Connect via `~/.kube/config` (override with `--kubeconfig`); fall back to in-cluster config when running in a pod.
- List pods matching the selector → spawn a streamer goroutine per pod/container.
- Watch for **new** pods matching the selector → spawn streamers when they go Ready.
- Each log line is tagged `[pod-name/container]` with a stable per-pod ANSI color.
- `--grep` filters lines by regex. `--json` emits `{ts, pod, container, line}` per line.
- Ctrl-C / SIGTERM → cancel the root context → all streamers exit → printer drains → exit 0.

## Architecture

```
            ┌──────────────────────────────────────────┐
            │                  main.go                 │
            │ flag.Parse → signal.NotifyContext → run  │
            └──────────────┬───────────────────────────┘
                           │ ctx
                ┌──────────▼───────────┐
                │     internal/k8s     │ kubeconfig + clientset
                └──────────┬───────────┘
                           │
                ┌──────────▼───────────┐
                │  internal/streamer   │
                │ - lists matching pods│
                │ - watches for new ones│
                │ - 1 goroutine per pod │
                │ - sends LogLine on   │
                │   shared chan         │
                └──────────┬───────────┘
                           │ chan LogLine
                ┌──────────▼───────────┐
                │  internal/printer    │
                │ - formats one line   │
                │ - color or --json    │
                │ - writes to os.Stdout│
                └──────────────────────┘
```

## Implementation order

Tackle it in stages — verify each before moving on.

### Stage 1 — boot the binary

```sh
cd 17-capstone
go build ./cmd/kubetail
./kubetail --help
```

`flag.Parse` should print usage. Already wired up in `cmd/kubetail/main.go`.

### Stage 2 — kubeconfig + clientset

`internal/k8s/client.go` is provided. It loads kubeconfig with the standard
discovery (explicit path → in-cluster → `~/.kube/config`) and returns a
`kubernetes.Interface`. Verify against a kind cluster:

```sh
kind create cluster --name kubetail-lab
kubectl apply -f testdata/loggy.yaml
./kubetail -n default -l app=loggy --once    # implement --once first to just list
```

### Stage 3 — single pod log stream

In `internal/streamer/streamer.go`, implement `streamPod` — open the log
stream via `clientset.CoreV1().Pods(ns).GetLogs(name, opts).Stream(ctx)`, scan
lines, send `LogLine` values to the shared channel, return on EOF or
ctx.Done(). Tests in `streamer_test.go` use the fake clientset.

### Stage 4 — fan-in multi-pod

Implement `Run` (the orchestrator):

1. List pods matching the selector.
2. For each Ready pod's containers, spawn a `streamPod` goroutine.
3. Drain the shared `LogLine` channel into the printer.
4. On `ctx.Done()`, all streamers exit; the orchestrator closes the channel
   when the last one finishes (use `sync.WaitGroup`).

Don't let the printer interleave partial lines from two pods — only complete
lines flow through the channel.

### Stage 5 — watch for new pods

Add a watcher goroutine using `clientset.CoreV1().Pods(ns).Watch(ctx, opts)`.
When you see an `ADDED` or `MODIFIED` event for a pod that's now Ready and
not already being streamed, spawn a streamer for it.

Track which pods you're already tailing in a `map[string]bool` guarded by a
`sync.Mutex`.

### Stage 6 — formatting + flags

Implement `printer.FormatLine` for human (color-coded prefix + line) and
`--json` modes.

Add `--grep` filtering in the streamer (cheaper than the printer — drop early).

## Make targets

```sh
make build       # go build ./cmd/kubetail → ./kubetail
make test        # unit tests (no cluster needed; uses fake clientset)
make e2e         # full kind-cluster end-to-end (requires kind + kubectl)
make lint        # gofmt + go vet + golangci-lint (if installed)
make clean
```

## Stretch goals

- TUI dashboard with `bubbletea` (split panes per pod, scroll).
- `--follow=false` snapshot mode (current logs once and exit).
- Save sessions to file with rotation.
- Deploy kubetail itself as an in-cluster `Job` with a minimal RBAC manifest
  (pods/log read in the target namespace).

## What you'll have built

A tool you'll actually run on the job. Stick it on your resume. When the
interview asks "tell me about something you've built in Go," you have your
answer: a multi-pod log tailer that auto-attaches, that you wrote as a learning
project covering goroutines, channels, context, client-go, and observability.
That's a real story.
