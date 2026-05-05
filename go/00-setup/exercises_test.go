package setup

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting()
	want := "go lab is alive"
	if got != want {
		t.Fatalf("Greeting() = %q, want %q", got, want)
	}
}
