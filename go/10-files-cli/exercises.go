package filescli

import (
	"flag"
	"io"
)

// ----- WordCount-from-reader -----

// WordCount reports lines, words, and bytes consumed from r.
//
//   - lines = number of newlines + (1 if last line is non-empty without trailing \n)
//     ...just kidding. Use bufio.Scanner with default ScanLines: lines = number
//     of times Scan() returned true.
//   - words = total whitespace-separated tokens across all lines.
//   - bytes = total bytes read.
//
// Examples:
//
//	WordCount("hello world\nfoo bar baz\n") -> 2, 5, 24
//	WordCount("")                            -> 0, 0, 0
//
// Hint: bufio.Scanner for lines. strings.Fields(line) for words. Sum len(line)+1
// per line you scan (the +1 is for the stripped \n) — but be careful for the
// final line if no trailing newline. Easier: count bytes by reading them all
// first with io.ReadAll, then process.
func WordCount(r io.Reader) (lines, words, bytes int, err error) {
	// TODO
	return 0, 0, 0, nil
}

// ----- Atomic write -----

// WriteAtomically writes data to path durably:
//
//  1. Write to path + ".tmp" (perm 0o644).
//  2. os.Rename to the final path.
//  3. On any failure during step 1, attempt to remove the .tmp.
//
// Returns nil on full success.
func WriteAtomically(path string, data []byte) error {
	// TODO — use os.WriteFile + os.Rename + os.Remove
	return nil
}

// ----- Flag parsing -----

// AppConfig is what we extract from CLI flags.
type AppConfig struct {
	Addr    string
	Workers int
	Debug   bool
}

// ParseFlags parses args (without the program name) using a *flag.FlagSet —
// NOT the global flag package — so the function is testable.
//
// Flags:
//
//	-addr string   listen address (default ":8080")
//	-workers int   concurrent workers (default 4)
//	-debug         debug logging (bool, default false)
//
// Return ("usage error" via err) if Parse fails. Use flag.ContinueOnError so
// failures return rather than os.Exit.
func ParseFlags(args []string) (AppConfig, error) {
	// TODO — fs := flag.NewFlagSet("app", flag.ContinueOnError)
	_ = flag.ContinueOnError
	return AppConfig{}, nil
}
