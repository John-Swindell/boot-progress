package solutions

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrDivByZero = errors.New("division by zero")
	ErrEmpty     = errors.New("empty slice")
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}

func ParseAndDouble(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parse and double: %w", err)
	}
	return 2 * n, nil
}

func MinMax(xs []int) (min, max int, err error) {
	if len(xs) == 0 {
		err = ErrEmpty
		return
	}
	min, max = xs[0], xs[0]
	for _, x := range xs[1:] {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	return
}

type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return "validation failed: " + e.Field
}

func Validate(name string) error {
	if name == "" {
		return &ValidationError{Field: "name"}
	}
	return nil
}

func Retry(attempts int, fn func() error) error {
	if attempts <= 0 {
		return errors.New("retry: attempts must be positive")
	}
	var last error
	for i := 0; i < attempts; i++ {
		if err := fn(); err != nil {
			last = err
			continue
		}
		return nil
	}
	return fmt.Errorf("retry exhausted: %w", last)
}
