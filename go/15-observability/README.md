# 15 — Observability

> **Goal:** wire up Prometheus metrics, expose pprof, and know what each metric
> type is for. Every prod Go service ships these. The capstone uses them.

## What you'll learn

- Prometheus metric types: Counter, Gauge, Histogram, Summary
- `prometheus.Registry` — explicit registry vs the global default
- Labels: when to use them, when they'll explode your cardinality
- `promhttp.HandlerFor` — exposing `/metrics`
- `net/http/pprof` — CPU/heap/goroutine profiles
- A taste of OpenTelemetry tracing (read-only — full setup is its own course)

## Coming from Python

| Python (`prometheus_client`) | Go (`client_golang`) | Note |
|---|---|---|
| `Counter('foo')` | `prometheus.NewCounter(prometheus.CounterOpts{Name: "foo"})` | Verbose but explicit. |
| `c.inc()` | `c.Inc()` | Same. |
| `c.labels(method='GET').inc()` | `c.WithLabelValues("GET").Inc()` | |
| `start_http_server(8000)` | `http.Handle("/metrics", promhttp.Handler()); http.ListenAndServe(...)` | You wire it into your existing server. |
| Default registry | `prometheus.DefaultRegisterer` (avoid in prod) | Prefer your own `prometheus.NewRegistry()` for testability. |

## Metric types — when to use each

- **Counter** — only increases. Requests, errors, bytes processed. Resets to 0 on process restart.
- **Gauge** — goes up *and* down. Connections in flight, queue depth, current goroutines.
- **Histogram** — distribution of observed values. Request latency, response sizes. Pre-allocated buckets.
- **Summary** — like Histogram but quantiles computed in-process. Less Prometheus-friendly; prefer Histogram unless you have a reason.

## Cardinality — the metric killer

Every unique label combination is a separate time series. **Never** use
unbounded labels (user IDs, request IDs, full URLs). Stick to small enumerated
sets: `method`, `status_code` (2xx/4xx/5xx — bucket if needed), `endpoint`
(static templated path, not the actual URL).

## pprof — free CPU/memory profiles

```go
import _ "net/http/pprof"

go func() { http.ListenAndServe(":6060", nil) }()
```

That blank import registers `/debug/pprof/...` handlers on the default mux.
Then:

```sh
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30   # CPU
go tool pprof http://localhost:6060/debug/pprof/heap                  # heap
go tool pprof http://localhost:6060/debug/pprof/goroutine             # goroutines
```

In an interactive session: `top`, `list <funcname>`, `web` (graphviz).

For prod, mount pprof on a **separate** port that isn't exposed publicly — it
leaks a lot of internals.

## Your turn

```sh
go test ./15-observability
```
