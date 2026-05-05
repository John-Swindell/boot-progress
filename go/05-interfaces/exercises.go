package interfaces

import "io"

// ----- Stringer -----

// Severity is a log level. Implement String() to make it satisfy fmt.Stringer.
type Severity int

const (
	SevDebug Severity = iota
	SevInfo
	SevWarn
	SevError
)

// String returns one of "DEBUG", "INFO", "WARN", "ERROR".
// Unknown values return "UNKNOWN".
func (s Severity) String() string {
	// TODO
	return ""
}

// ----- Notifier -----

// Notifier sends a message somewhere. Anything with a Notify method satisfies
// it. Used to decouple "send a message" from where it's sent.
type Notifier interface {
	Notify(msg string) error
}

// SliceNotifier records every notification it receives. Useful for tests.
//
// Hint: pointer receiver — you're mutating Sent.
type SliceNotifier struct {
	Sent []string
}

// Notify appends msg to s.Sent and returns nil.
func (s *SliceNotifier) Notify(msg string) error {
	// TODO
	return nil
}

// FailingNotifier always returns the supplied Err.
type FailingNotifier struct {
	Err error
}

// Notify returns f.Err for any input.
func (f *FailingNotifier) Notify(msg string) error {
	// TODO
	return nil
}

// NotifyAll calls Notify(msg) on every notifier. It returns the FIRST non-nil
// error encountered, but only AFTER attempting every notifier (so logging and
// metrics still fire even if email fails).
//
// Returns nil if all succeed.
func NotifyAll(ns []Notifier, msg string) error {
	// TODO
	return nil
}

// ----- Type switch -----

// Describe returns one of "int", "string", "slice", "map", "nil", "unknown".
//   - any nil interface value (including typed nils that come in as nil) -> "nil"
//   - int, int32, int64 -> "int"
//   - string -> "string"
//   - any []T -> "slice"   (use type switch with []int, []string, []any cases)
//   - map[K]V where K is comparable -> "map" (use cases for map[string]string, map[string]int, map[string]any)
//   - everything else -> "unknown"
func Describe(x any) string {
	// TODO
	return ""
}

// ----- io.Writer -----

// WriteGreeting writes "hello, <name>\n" to w. It returns the number of bytes
// written and any error from w.
//
// Hint: fmt.Fprintf returns (n int, err error) and writes to any io.Writer.
func WriteGreeting(w io.Writer, name string) (int, error) {
	// TODO
	return 0, nil
}
