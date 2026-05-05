package observ

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func TestMetricsAndHandler(t *testing.T) {
	reg := prometheus.NewRegistry()
	var registerer prometheus.Registerer = reg
	m := NewMetrics(&registerer)
	if m == nil || m.CheckTotal == nil || m.CheckSeconds == nil {
		t.Fatalf("NewMetrics returned %+v", m)
	}

	m.RecordCheck("github", true, 50*time.Millisecond)
	m.RecordCheck("github", true, 30*time.Millisecond)
	m.RecordCheck("github", false, 100*time.Millisecond)
	m.RecordCheck("sshd", true, 1*time.Millisecond)

	srv := httptest.NewServer(MetricsHandler(reg))
	defer srv.Close()

	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Fatalf("GET: %v", err)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	out := string(body)

	mustContain := []string{
		`check_total{name="github",result="ok"} 2`,
		`check_total{name="github",result="fail"} 1`,
		`check_total{name="sshd",result="ok"} 1`,
		`check_seconds_count{name="github"} 3`,
		`check_seconds_count{name="sshd"} 1`,
	}
	for _, want := range mustContain {
		if !strings.Contains(out, want) {
			t.Errorf("metrics output missing %q\n--- output ---\n%s", want, out)
		}
	}
}
