package solutions

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	CheckTotal   *prometheus.CounterVec
	CheckSeconds *prometheus.HistogramVec
}

func NewMetrics(reg *prometheus.Registerer) *Metrics {
	m := &Metrics{
		CheckTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "check_total",
			Help: "number of checks executed",
		}, []string{"name", "result"}),
		CheckSeconds: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "check_seconds",
			Help:    "duration of checks in seconds",
			Buckets: prometheus.DefBuckets,
		}, []string{"name"}),
	}
	(*reg).MustRegister(m.CheckTotal, m.CheckSeconds)
	return m
}

func (m *Metrics) RecordCheck(name string, success bool, dur time.Duration) {
	result := "ok"
	if !success {
		result = "fail"
	}
	m.CheckTotal.WithLabelValues(name, result).Inc()
	m.CheckSeconds.WithLabelValues(name).Observe(dur.Seconds())
}

func MetricsHandler(reg *prometheus.Registry) http.Handler {
	return promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
}
