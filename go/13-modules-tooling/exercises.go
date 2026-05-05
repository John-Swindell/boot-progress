package tooling

// Version, Commit, and Date are intended to be overwritten at link time using
// `-ldflags "-X example.com/golab/13-modules-tooling.Version=v1.2.3"`.
//
// They must be package-level `var`s of type string for -X to work
// (-X cannot set consts).
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

// BuildInfo returns a one-line human-readable build identifier of the form:
//
//	"<Version> (<Commit>) built <Date>"
//
// Example: "v1.2.3 (abc1234) built 2024-05-01T12:00:00Z"
//
// TODO: implement using the package vars above.
func BuildInfo() string {
	// TODO
	return ""
}
