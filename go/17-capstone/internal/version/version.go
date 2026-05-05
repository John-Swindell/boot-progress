// Package version exports build metadata that's stamped at link time via
// -ldflags "-X .../version.Version=...".
package version

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

// String returns "<Version> (<Commit>) built <Date>".
func String() string {
	return Version + " (" + Commit + ") built " + Date
}
