package stdlib

import (
	"strings"
	"testing"
	"time"
)

func TestFormatBytes(t *testing.T) {
	cases := []struct {
		in   int64
		want string
	}{
		{-1, "invalid"},
		{0, "0 B"},
		{1023, "1023 B"},
		{1024, "1.0 KiB"},
		{1536, "1.5 KiB"},
		{1048576, "1.0 MiB"},
		{2 * 1024 * 1024 * 1024, "2.0 GiB"},
	}
	for _, tc := range cases {
		if got := FormatBytes(tc.in); got != tc.want {
			t.Errorf("FormatBytes(%d) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestExpoBackoff(t *testing.T) {
	base := 100 * time.Millisecond
	max := 5 * time.Second
	cases := []struct {
		attempt int
		want    time.Duration
	}{
		{-1, 0},
		{0, base},
		{1, 200 * time.Millisecond},
		{3, 800 * time.Millisecond},
		{20, max},
	}
	for _, tc := range cases {
		if got := ExpoBackoff(tc.attempt, base, max); got != tc.want {
			t.Errorf("ExpoBackoff(%d) = %v, want %v", tc.attempt, got, tc.want)
		}
	}
}

func TestEncodeUser(t *testing.T) {
	created, _ := time.Parse(time.RFC3339, "2024-01-02T03:04:05Z")
	u := User{Name: "alice", Email: "a@x.io", Created: created, Notes: "secret"}
	s, err := EncodeUser(u)
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if !strings.Contains(s, `"name":"alice"`) {
		t.Errorf("missing name: %s", s)
	}
	if !strings.Contains(s, `"email":"a@x.io"`) {
		t.Errorf("missing email: %s", s)
	}
	if !strings.Contains(s, `"created_at":"2024-01-02T03:04:05Z"`) {
		t.Errorf("missing created_at: %s", s)
	}
	if strings.Contains(s, "secret") || strings.Contains(s, "Notes") {
		t.Errorf("Notes leaked: %s", s)
	}

	// omitempty
	u2 := User{Name: "bob", Created: created}
	s2, _ := EncodeUser(u2)
	if strings.Contains(s2, "email") {
		t.Errorf("expected email omitted: %s", s2)
	}
}

func TestDecodeUser(t *testing.T) {
	s := `{"name":"alice","email":"a@x.io","created_at":"2024-01-02T03:04:05Z"}`
	u, err := DecodeUser(s)
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if u.Name != "alice" || u.Email != "a@x.io" {
		t.Errorf("got %+v", u)
	}
	if !u.Created.Equal(time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)) {
		t.Errorf("Created = %v", u.Created)
	}
}

func TestParseLogLine(t *testing.T) {
	e, err := ParseLogLine("2024-05-01T12:34:56Z INFO server started on :8080")
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if e.Level != "INFO" || e.Message != "server started on :8080" {
		t.Errorf("got %+v", e)
	}
	if !e.Time.Equal(time.Date(2024, 5, 1, 12, 34, 56, 0, time.UTC)) {
		t.Errorf("Time = %v", e.Time)
	}

	if _, err := ParseLogLine("garbage"); err == nil {
		t.Error("expected error for malformed line")
	}
	if _, err := ParseLogLine("notatimestamp INFO hi"); err == nil {
		t.Error("expected error for bad timestamp")
	}
}
