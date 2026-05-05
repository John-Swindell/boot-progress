# 11 — HTTP Clients & Servers

> **Goal:** write servers and clients with `net/http`. No frameworks. The
> stdlib is enough for everything you'll do as a platform engineer.

## What you'll learn

- `http.HandlerFunc`, `http.ServeMux`, the `Handler` interface
- Reading queries (`r.URL.Query()`) and JSON request bodies
- Writing JSON responses (`w.Header().Set` + `json.NewEncoder(w).Encode`)
- Middleware (a function that takes a Handler and returns a Handler)
- Client side: `http.NewRequestWithContext`, `http.DefaultClient.Do`, JSON decode
- Graceful shutdown with `srv.Shutdown(ctx)` (preview, used heavily in capstone)
- `net/http/httptest` for testing

## Coming from Python

| Python (Flask / FastAPI) | Go | Note |
|---|---|---|
| `@app.route("/x")` | `mux.HandleFunc("/x", handler)` | No decorators — just register the func. |
| `request.args.get("name")` | `r.URL.Query().Get("name")` | Returns "" if missing — no `request.args["name"]` exception. |
| `return jsonify(d)` | `json.NewEncoder(w).Encode(d)` | Set `Content-Type` header first. |
| `before_request` | Middleware: `func(http.Handler) http.Handler` | Composable. Wrap the mux. |
| `requests.get(url)` | `http.Get(url)` (or `Do(req)` with context) | Always pass a context. |

## The Handler interface

```go
type Handler interface {
    ServeHTTP(w http.ResponseWriter, r *http.Request)
}
```

That's it. Anything with `ServeHTTP` is a handler. `http.HandlerFunc` is an
adapter so a plain `func(w, r)` satisfies the interface.

## Middleware pattern

```go
func WithLogging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        rec := &statusRecorder{ResponseWriter: w, status: 200}
        next.ServeHTTP(rec, r)
        slog.Info("http", "method", r.Method, "path", r.URL.Path,
            "status", rec.status, "dur", time.Since(start))
    })
}
```

To capture the status, wrap `ResponseWriter` and intercept `WriteHeader`. The
exercise walks you through this.

## Graceful shutdown

```go
srv := &http.Server{Addr: ":8080", Handler: mux}
go func() { srv.ListenAndServe() }()

<-ctx.Done()                         // SIGTERM came in (chapter 14)
ctxShut, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
srv.Shutdown(ctxShut)                // waits for in-flight requests
```

`Shutdown` stops accepting new connections, lets in-flight finish, returns when
done (or when its context fires).

## Testing with httptest

```go
srv := httptest.NewServer(handler)
defer srv.Close()
resp, err := http.Get(srv.URL + "/x")
```

Spins up a real HTTP server on a random port. The fastest, most realistic way
to test handlers — no mocking required.

## Your turn

```sh
go test ./11-http
```
