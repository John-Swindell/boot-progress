package printer

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestPrintJSON(t *testing.T) {
	var buf bytes.Buffer
	p := New(&buf, Options{JSON: true})
	ts := time.Date(2024, 5, 1, 12, 0, 0, 0, time.UTC)
	if err := p.Print(LogLine{Time: ts, Pod: "p1", Container: "c1", Text: "hello"}); err != nil {
		t.Fatalf("Print: %v", err)
	}

	out := strings.TrimRight(buf.String(), "\n")
	var got map[string]string
	if err := json.Unmarshal([]byte(out), &got); err != nil {
		t.Fatalf("not JSON: %v\n%s", err, out)
	}
	if got["pod"] != "p1" || got["container"] != "c1" || got["line"] != "hello" {
		t.Errorf("got %v", got)
	}
	if got["ts"] == "" {
		t.Errorf("missing ts")
	}
}

func TestPrintPretty(t *testing.T) {
	t.Run("with color", func(t *testing.T) {
		var buf bytes.Buffer
		p := New(&buf, Options{})
		_ = p.Print(LogLine{Pod: "p1", Container: "c1", Text: "hi"})
		s := buf.String()
		if !strings.Contains(s, "[p1/c1]") || !strings.Contains(s, "hi") {
			t.Errorf("got %q", s)
		}
		if !strings.Contains(s, "\x1b[") {
			t.Errorf("expected ANSI escape, got %q", s)
		}
		if !strings.HasSuffix(s, "\n") {
			t.Errorf("missing newline: %q", s)
		}
	})

	t.Run("no color", func(t *testing.T) {
		var buf bytes.Buffer
		p := New(&buf, Options{NoColor: true})
		_ = p.Print(LogLine{Pod: "p1", Container: "c1", Text: "hi"})
		s := buf.String()
		if strings.Contains(s, "\x1b[") {
			t.Errorf("unexpected ANSI: %q", s)
		}
		if s != "[p1/c1] hi\n" {
			t.Errorf("got %q, want %q", s, "[p1/c1] hi\n")
		}
	})
}

func TestPrintConcurrent(t *testing.T) {
	var buf bytes.Buffer
	p := New(&buf, Options{NoColor: true})
	done := make(chan struct{})
	for i := 0; i < 4; i++ {
		go func() {
			defer func() { done <- struct{}{} }()
			for j := 0; j < 100; j++ {
				p.Print(LogLine{Pod: "p", Container: "c", Text: "line"})
			}
		}()
	}
	for i := 0; i < 4; i++ {
		<-done
	}
	for _, line := range strings.Split(strings.TrimRight(buf.String(), "\n"), "\n") {
		if line != "[p/c] line" {
			t.Fatalf("torn line: %q", line)
		}
	}
}
