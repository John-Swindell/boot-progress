package observ

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Metrics bundles the Prometheus instruments for a "checker" service.
//
// CheckTotal:   counter labeled by (name, result) where result is "ok" or "fail"
// CheckSeconds: histogram labeled by (name) — observe seconds, default buckets
type Metrics struct {
	CheckTotal   *prometheus.CounterVec
	CheckSeconds *prometheus.HistogramVec
}

// NewMetrics constructs a Metrics, registering both vectors on reg.
//
//   - CheckTotal name:   "check_total"
//     help:   "number of checks executed"
//     labels: "name", "result"
//
//   - CheckSeconds name: "check_seconds"
//     help: "duration of checks in seconds"
//     labels: "name"
//     buckets: prometheus.DefBuckets
//
// Use prometheus.NewCounterVec / NewHistogramVec, then reg.MustRegister.
func NewMetrics(reg *prometheus.Registerer) *Metrics {
	// TODO
	return nil
}

// RecordCheck records one execution: increments CheckTotal{name,result}, where
// result is "ok" if success else "fail", and observes dur in CheckSeconds{name}.
func (m *Metrics) RecordCheck(name string, success bool, dur time.Duration) {
	// TODO
}

// MetricsHandler returns an http.Handler exposing reg in Prometheus text
// format. Use promhttp.HandlerFor with promhttp.HandlerOpts{}.
//
// Hint: import "github.com/prometheus/client_golang/prometheus/promhttp"
func MetricsHandler(reg *prometheus.Registry) http.Handler {
	// TODO
	return nil
}
