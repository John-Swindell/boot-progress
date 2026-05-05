package filescli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWordCount(t *testing.T) {
	cases := []struct {
		name    string
		in      string
		l, w, b int
	}{
		{"empty", "", 0, 0, 0},
		{"two lines", "hello world\nfoo bar baz\n", 2, 5, 24},
		{"no trailing newline", "one two\nthree", 2, 3, 13},
		{"only whitespace", "   \n\t\n", 2, 0, 6},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			l, w, b, err := WordCount(strings.NewReader(tc.in))
			if err != nil {
				t.Fatalf("err = %v", err)
			}
			if l != tc.l || w != tc.w || b != tc.b {
				t.Errorf("got (%d,%d,%d), want (%d,%d,%d)", l, w, b, tc.l, tc.w, tc.b)
			}
		})
	}
}

func TestWriteAtomically(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.txt")

	if err := WriteAtomically(path, []byte("hello")); err != nil {
		t.Fatalf("err = %v", err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if !bytes.Equal(got, []byte("hello")) {
		t.Errorf("got %q, want %q", got, "hello")
	}
	// .tmp must be gone after success
	if _, err := os.Stat(path + ".tmp"); !os.IsNotExist(err) {
		t.Errorf(".tmp left behind: %v", err)
	}

	// Overwrite still works
	if err := WriteAtomically(path, []byte("world")); err != nil {
		t.Fatalf("overwrite: %v", err)
	}
	got, _ = os.ReadFile(path)
	if !bytes.Equal(got, []byte("world")) {
		t.Errorf("after overwrite got %q, want %q", got, "world")
	}
}

func TestParseFlags(t *testing.T) {
	t.Run("defaults", func(t *testing.T) {
		cfg, err := ParseFlags([]string{})
		if err != nil {
			t.Fatalf("err = %v", err)
		}
		if cfg.Addr != ":8080" || cfg.Workers != 4 || cfg.Debug != false {
			t.Errorf("got %+v", cfg)
		}
	})
	t.Run("override all", func(t *testing.T) {
		cfg, err := ParseFlags([]string{"-addr=:9090", "-workers=8", "-debug"})
		if err != nil {
			t.Fatalf("err = %v", err)
		}
		if cfg.Addr != ":9090" || cfg.Workers != 8 || cfg.Debug != true {
			t.Errorf("got %+v", cfg)
		}
	})
	t.Run("unknown flag errors", func(t *testing.T) {
		if _, err := ParseFlags([]string{"-nope"}); err == nil {
			t.Error("expected error for unknown flag")
		}
	})
}
