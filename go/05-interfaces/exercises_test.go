package interfaces

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

func TestSeverityString(t *testing.T) {
	cases := []struct {
		s    Severity
		want string
	}{
		{SevDebug, "DEBUG"},
		{SevInfo, "INFO"},
		{SevWarn, "WARN"},
		{SevError, "ERROR"},
		{Severity(99), "UNKNOWN"},
	}
	for _, tc := range cases {
		if got := tc.s.String(); got != tc.want {
			t.Errorf("%d.String() = %q, want %q", int(tc.s), got, tc.want)
		}
		// Also verify it satisfies fmt.Stringer via fmt.Sprint.
		if got := fmt.Sprint(tc.s); got != tc.want {
			t.Errorf("fmt.Sprint(%d) = %q, want %q", int(tc.s), got, tc.want)
		}
	}
}

func TestSliceNotifier(t *testing.T) {
	var n Notifier = &SliceNotifier{}
	if err := n.Notify("hello"); err != nil {
		t.Errorf("Notify err = %v", err)
	}
	sn := n.(*SliceNotifier)
	if len(sn.Sent) != 1 || sn.Sent[0] != "hello" {
		t.Errorf("Sent = %v, want [hello]", sn.Sent)
	}
}

func TestNotifyAll(t *testing.T) {
	good := &SliceNotifier{}
	bad := &FailingNotifier{Err: errors.New("smtp down")}
	other := &SliceNotifier{}

	err := NotifyAll([]Notifier{good, bad, other}, "ping")
	if err == nil || err.Error() != "smtp down" {
		t.Errorf("err = %v, want smtp down", err)
	}
	// every notifier was still attempted
	if len(good.Sent) != 1 || len(other.Sent) != 1 {
		t.Errorf("not all notifiers attempted; good=%v other=%v", good.Sent, other.Sent)
	}

	if err := NotifyAll([]Notifier{good, other}, "x"); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestDescribe(t *testing.T) {
	cases := []struct {
		in   any
		want string
	}{
		{nil, "nil"},
		{42, "int"},
		{int64(7), "int"},
		{"hi", "string"},
		{[]int{1, 2}, "slice"},
		{[]string{"a"}, "slice"},
		{map[string]int{"x": 1}, "map"},
		{3.14, "unknown"},
		{struct{}{}, "unknown"},
	}
	for _, tc := range cases {
		if got := Describe(tc.in); got != tc.want {
			t.Errorf("Describe(%v) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestWriteGreeting(t *testing.T) {
	var buf bytes.Buffer
	n, err := WriteGreeting(&buf, "world")
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	want := "hello, world\n"
	if buf.String() != want {
		t.Errorf("buf = %q, want %q", buf.String(), want)
	}
	if n != len(want) {
		t.Errorf("n = %d, want %d", n, len(want))
	}
}
