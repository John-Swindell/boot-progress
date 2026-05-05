package solutions

import "fmt"

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

func BuildInfo() string {
	return fmt.Sprintf("%s (%s) built %s", Version, Commit, Date)
}
