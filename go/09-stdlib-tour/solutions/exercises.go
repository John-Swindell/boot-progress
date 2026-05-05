package solutions

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func FormatBytes(n int64) string {
	if n < 0 {
		return "invalid"
	}
	if n < 1024 {
		return fmt.Sprintf("%d B", n)
	}
	units := []string{"KiB", "MiB", "GiB", "TiB"}
	val := float64(n) / 1024
	i := 0
	for val >= 1024 && i < len(units)-1 {
		val /= 1024
		i++
	}
	return fmt.Sprintf("%.1f %s", val, units[i])
}

func ExpoBackoff(attempt int, base, max time.Duration) time.Duration {
	if attempt < 0 || base <= 0 || max <= 0 {
		return 0
	}
	d := base << attempt
	if d <= 0 || d > max {
		return max
	}
	return d
}

type User struct {
	Name    string    `json:"name"`
	Email   string    `json:"email,omitempty"`
	Created time.Time `json:"created_at"`
	Notes   string    `json:"-"`
}

func EncodeUser(u User) (string, error) {
	b, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func DecodeUser(s string) (User, error) {
	var u User
	err := json.Unmarshal([]byte(s), &u)
	return u, err
}

type LogEntry struct {
	Time    time.Time
	Level   string
	Message string
}

func ParseLogLine(line string) (LogEntry, error) {
	parts := strings.SplitN(line, " ", 3)
	if len(parts) != 3 {
		return LogEntry{}, fmt.Errorf("malformed log line: %q", line)
	}
	t, err := time.Parse(time.RFC3339, parts[0])
	if err != nil {
		return LogEntry{}, fmt.Errorf("bad timestamp: %w", err)
	}
	return LogEntry{Time: t, Level: parts[1], Message: parts[2]}, nil
}
