package httplab

import (
	"context"
	"net/http"
)

// HelloHandler responds to GET requests with "hello, <name>\n" where <name>
// comes from the "name" query parameter. If absent or empty, default to "world".
//
// Set Content-Type: text/plain; charset=utf-8.
// On non-GET methods, return 405 Method Not Allowed.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// HealthHandler responds 200 OK with JSON body `{"status":"ok"}`.
// Set Content-Type: application/json.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// statusRecorder is provided to help with the LoggingMiddleware exercise.
// It captures the response status without changing observable behavior.
type statusRecorder struct {
	http.ResponseWriter
	Status int
}

func (s *statusRecorder) WriteHeader(code int) {
	s.Status = code
	s.ResponseWriter.WriteHeader(code)
}

// LoggingMiddleware wraps next so that the resulting Handler:
//  1. Records the response status (default 200 if Write happens without WriteHeader).
//  2. After calling next, sets the response header "X-Status-Logged" to the
//     captured status as a decimal string (e.g. "200", "404").
//
// The header must be set BEFORE the response is sent. Hint: set X-Status-Logged
// AFTER next.ServeHTTP returns won't work because headers are flushed on first
// Write. Workaround for this exercise: we'll *also* re-emit the header inside
// statusRecorder by inspecting after — to keep this testable, set the header
// once you're done by writing it on the recorder. Easier exercise framing:
// instead of a header, append a slog-style line to the *Logs slice (passed in).
//
// To keep it simple and testable, this version of LoggingMiddleware appends
// "<METHOD> <PATH> <STATUS>" to the Logs slice (pointer) AFTER next runs.
type LogSink struct {
	Lines []string
}

func WithLogging(sink *LogSink, next http.Handler) http.Handler {
	// TODO — return an http.HandlerFunc that:
	//  1. wraps w in &statusRecorder{ResponseWriter: w, Status: 200}
	//  2. calls next.ServeHTTP(rec, r)
	//  3. appends fmt.Sprintf("%s %s %d", r.Method, r.URL.Path, rec.Status) to sink.Lines
	return next
}

// FetchJSON GETs url with the supplied context and decodes the JSON response
// into target (a pointer). Returns error on non-2xx status (with the status
// in the message) or any I/O / decode error.
//
// Always pass the request through the supplied http.Client. nil client means
// http.DefaultClient.
func FetchJSON(ctx context.Context, client *http.Client, url string, target any) error {
	// TODO
	return nil
}
