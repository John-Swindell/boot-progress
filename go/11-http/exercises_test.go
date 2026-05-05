package httplab

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	t.Run("default world", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		rr := httptest.NewRecorder()
		HelloHandler(rr, req)
		if rr.Code != 200 {
			t.Errorf("status = %d, want 200", rr.Code)
		}
		if !strings.Contains(rr.Body.String(), "hello, world") {
			t.Errorf("body = %q, want hello, world", rr.Body.String())
		}
		if ct := rr.Header().Get("Content-Type"); !strings.HasPrefix(ct, "text/plain") {
			t.Errorf("Content-Type = %q", ct)
		}
	})
	t.Run("with name", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/?name=alice", nil)
		rr := httptest.NewRecorder()
		HelloHandler(rr, req)
		if !strings.Contains(rr.Body.String(), "hello, alice") {
			t.Errorf("body = %q", rr.Body.String())
		}
	})
	t.Run("rejects POST", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/", nil)
		rr := httptest.NewRecorder()
		HelloHandler(rr, req)
		if rr.Code != http.StatusMethodNotAllowed {
			t.Errorf("status = %d, want 405", rr.Code)
		}
	})
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/healthz", nil)
	rr := httptest.NewRecorder()
	HealthHandler(rr, req)
	if rr.Code != 200 {
		t.Errorf("status = %d", rr.Code)
	}
	var body map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &body); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if body["status"] != "ok" {
		t.Errorf("body = %v", body)
	}
}

func TestWithLogging(t *testing.T) {
	sink := &LogSink{}
	mux := http.NewServeMux()
	mux.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
		w.Write([]byte("created"))
	})
	mux.HandleFunc("/teapot", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(418)
	})
	wrapped := WithLogging(sink, mux)

	srv := httptest.NewServer(wrapped)
	defer srv.Close()

	if _, err := http.Get(srv.URL + "/ok"); err != nil {
		t.Fatalf("get: %v", err)
	}
	if _, err := http.Get(srv.URL + "/teapot"); err != nil {
		t.Fatalf("get: %v", err)
	}

	if len(sink.Lines) != 2 {
		t.Fatalf("Lines = %v", sink.Lines)
	}
	if !strings.Contains(sink.Lines[0], "GET /ok 201") {
		t.Errorf("[0] = %q", sink.Lines[0])
	}
	if !strings.Contains(sink.Lines[1], "GET /teapot 418") {
		t.Errorf("[1] = %q", sink.Lines[1])
	}
}

func TestFetchJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/ok":
			io.WriteString(w, `{"name":"alice","age":30}`)
		case "/bad":
			http.Error(w, "nope", http.StatusInternalServerError)
		}
	}))
	defer srv.Close()

	t.Run("success", func(t *testing.T) {
		var out struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		err := FetchJSON(context.Background(), srv.Client(), srv.URL+"/ok", &out)
		if err != nil {
			t.Fatalf("err = %v", err)
		}
		if out.Name != "alice" || out.Age != 30 {
			t.Errorf("got %+v", out)
		}
	})

	t.Run("non-2xx", func(t *testing.T) {
		var out map[string]any
		err := FetchJSON(context.Background(), srv.Client(), srv.URL+"/bad", &out)
		if err == nil {
			t.Error("expected error")
		}
	})
}
