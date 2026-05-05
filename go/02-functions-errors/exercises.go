package errs

import "errors"

// ErrDivByZero is returned by Divide when b == 0.
// Callers should be able to detect it with errors.Is(err, ErrDivByZero).
var ErrDivByZero = errors.New("division by zero")

// Divide returns a/b, or (0, ErrDivByZero) when b == 0.
func Divide(a, b float64) (float64, error) {
	// TODO
	return 0, nil
}

// ParseAndDouble parses s as a base-10 integer and returns 2*n.
// On parse failure it returns an error wrapping the underlying strconv error
// with the prefix "parse and double: " (use fmt.Errorf with %w).
//
// Hint: import "strconv" and "fmt".
func ParseAndDouble(s string) (int, error) {
	// TODO
	return 0, nil
}

// MinMax returns the smallest and largest element of xs.
// For an empty slice it returns (0, 0, ErrEmpty).
var ErrEmpty = errors.New("empty slice")

func MinMax(xs []int) (min, max int, err error) {
	// TODO — use named returns; one pass; no sort.
	return
}

// ValidationError is a structured error reporting which field failed.
// It must implement the error interface.
type ValidationError struct {
	Field string
}

// Error makes *ValidationError satisfy the error interface.
//
// TODO: replace this stub. Return the string
//
//	"validation failed: " + e.Field
func (e *ValidationError) Error() string {
	// TODO
	return ""
}

// Validate returns nil if name is non-empty, else *ValidationError{Field: "name"}.
// Caller will use errors.As to recover the struct.
func Validate(name string) error {
	// TODO
	return nil
}

// Retry calls fn up to attempts times. It returns nil as soon as fn returns nil.
// If every attempt fails, it returns the LAST error wrapped with the prefix
// "retry exhausted: " (using fmt.Errorf %w).
//
// attempts <= 0 → return errors.New("retry: attempts must be positive") without
// calling fn.
func Retry(attempts int, fn func() error) error {
	// TODO
	return nil
}
