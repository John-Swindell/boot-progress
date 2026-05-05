package solutions

import (
	"fmt"
	"io"
)

type Severity int

const (
	SevDebug Severity = iota
	SevInfo
	SevWarn
	SevError
)

func (s Severity) String() string {
	switch s {
	case SevDebug:
		return "DEBUG"
	case SevInfo:
		return "INFO"
	case SevWarn:
		return "WARN"
	case SevError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

type Notifier interface {
	Notify(msg string) error
}

type SliceNotifier struct{ Sent []string }

func (s *SliceNotifier) Notify(msg string) error {
	s.Sent = append(s.Sent, msg)
	return nil
}

type FailingNotifier struct{ Err error }

func (f *FailingNotifier) Notify(msg string) error { return f.Err }

func NotifyAll(ns []Notifier, msg string) error {
	var firstErr error
	for _, n := range ns {
		if err := n.Notify(msg); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

func Describe(x any) string {
	if x == nil {
		return "nil"
	}
	switch x.(type) {
	case int, int32, int64:
		return "int"
	case string:
		return "string"
	case []int, []string, []any:
		return "slice"
	case map[string]string, map[string]int, map[string]any:
		return "map"
	default:
		return "unknown"
	}
}

func WriteGreeting(w io.Writer, name string) (int, error) {
	return fmt.Fprintf(w, "hello, %s\n", name)
}
