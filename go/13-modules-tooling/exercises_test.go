package tooling

import (
	"strings"
	"testing"
)

func TestBuildInfoDefaults(t *testing.T) {
	got := BuildInfo()
	want := "dev (none) built unknown"
	if got != want {
		t.Errorf("BuildInfo() = %q, want %q", got, want)
	}
}

func TestBuildInfoWithOverride(t *testing.T) {
	saveV, saveC, saveD := Version, Commit, Date
	defer func() { Version, Commit, Date = saveV, saveC, saveD }()

	Version = "v1.2.3"
	Commit = "abc1234"
	Date = "2024-05-01T12:00:00Z"

	got := BuildInfo()
	if !strings.Contains(got, "v1.2.3") || !strings.Contains(got, "abc1234") {
		t.Errorf("BuildInfo() = %q", got)
	}
}
