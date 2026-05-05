package stdlib

import "time"

// ----- FormatBytes -----

// FormatBytes returns a human-friendly size string using BINARY (1024) units.
// Examples:
//
//	0           -> "0 B"
//	1023        -> "1023 B"
//	1024        -> "1.0 KiB"
//	1536        -> "1.5 KiB"
//	1048576     -> "1.0 MiB"
//	2147483648  -> "2.0 GiB"
//
// Use units B, KiB, MiB, GiB, TiB. Negative inputs return "invalid".
//
// Hint: fmt.Sprintf with "%.1f %s".
func FormatBytes(n int64) string {
	// TODO
	return ""
}

// ----- ExpoBackoff -----

// ExpoBackoff returns base * 2^attempt, capped at max.
// attempt < 0 returns 0. base or max <= 0 returns 0.
//
//	ExpoBackoff(0, 100ms, 5s) -> 100ms
//	ExpoBackoff(3, 100ms, 5s) -> 800ms
//	ExpoBackoff(20, 100ms, 5s) -> 5s   (capped)
func ExpoBackoff(attempt int, base, max time.Duration) time.Duration {
	// TODO
	return 0
}

// ----- JSON -----

// User is the JSON shape we'll marshal/unmarshal.
// Field tags must produce JSON like:
//
//	{"name":"alice","email":"a@x.io","created_at":"2024-01-02T03:04:05Z"}
//
// The Email field must be omitted entirely when empty.
// The internal Notes field must NEVER appear in JSON.
type User struct {
	Name    string    // TODO: json:"name"
	Email   string    // TODO: json:"email,omitempty"
	Created time.Time // TODO: json:"created_at"
	Notes   string    // TODO: json:"-"
}

// EncodeUser returns the JSON representation of u as a string.
func EncodeUser(u User) (string, error) {
	// TODO — use encoding/json
	return "", nil
}

// DecodeUser parses s into a User.
func DecodeUser(s string) (User, error) {
	// TODO
	return User{}, nil
}

// ----- ParseLogLine -----

// LogEntry is one parsed structured-log line.
type LogEntry struct {
	Time    time.Time
	Level   string // "INFO", "WARN", "ERROR", etc.
	Message string
}

// ParseLogLine parses lines of the form:
//
//	2024-05-01T12:34:56Z INFO server started on :8080
//
// (RFC3339 timestamp, single space, level token, single space, free-form message)
//
// Returns an error for any malformed input.
//
// Hint: strings.SplitN(s, " ", 3); time.Parse(time.RFC3339, parts[0]).
func ParseLogLine(line string) (LogEntry, error) {
	// TODO
	return LogEntry{}, nil
}
