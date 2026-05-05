package errs

import (
	"errors"
	"strconv"
	"testing"
)

func TestDivide(t *testing.T) {
	got, err := Divide(10, 2)
	if err != nil || got != 5 {
		t.Errorf("Divide(10,2) = (%v, %v), want (5, nil)", got, err)
	}
	_, err = Divide(1, 0)
	if !errors.Is(err, ErrDivByZero) {
		t.Errorf("Divide(1,0) err = %v, want errors.Is ErrDivByZero", err)
	}
}

func TestParseAndDouble(t *testing.T) {
	got, err := ParseAndDouble("21")
	if err != nil || got != 42 {
		t.Errorf("ParseAndDouble(21) = (%v, %v), want (42, nil)", got, err)
	}
	_, err = ParseAndDouble("nope")
	if err == nil {
		t.Fatal("expected error for non-numeric input")
	}
	var nerr *strconv.NumError
	if !errors.As(err, &nerr) {
		t.Errorf("expected wrapped *strconv.NumError, got %T: %v", err, err)
	}
}

func TestMinMax(t *testing.T) {
	mn, mx, err := MinMax([]int{3, 1, 4, 1, 5, 9, 2, 6})
	if err != nil || mn != 1 || mx != 9 {
		t.Errorf("MinMax = (%d,%d,%v), want (1,9,nil)", mn, mx, err)
	}
	_, _, err = MinMax(nil)
	if !errors.Is(err, ErrEmpty) {
		t.Errorf("MinMax(nil) err = %v, want ErrEmpty", err)
	}
}

func TestValidate(t *testing.T) {
	if err := Validate("alice"); err != nil {
		t.Errorf("Validate(alice) = %v, want nil", err)
	}
	err := Validate("")
	var verr *ValidationError
	if !errors.As(err, &verr) {
		t.Fatalf("Validate(\"\") err = %v, want *ValidationError", err)
	}
	if verr.Field != "name" {
		t.Errorf("Field = %q, want %q", verr.Field, "name")
	}
	if got := err.Error(); got != "validation failed: name" {
		t.Errorf("Error() = %q, want %q", got, "validation failed: name")
	}
}

func TestRetry(t *testing.T) {
	t.Run("succeeds eventually", func(t *testing.T) {
		calls := 0
		err := Retry(3, func() error {
			calls++
			if calls < 2 {
				return errors.New("fail")
			}
			return nil
		})
		if err != nil {
			t.Errorf("err = %v, want nil", err)
		}
		if calls != 2 {
			t.Errorf("calls = %d, want 2", calls)
		}
	})

	t.Run("exhausts and wraps last error", func(t *testing.T) {
		boom := errors.New("boom")
		err := Retry(3, func() error { return boom })
		if !errors.Is(err, boom) {
			t.Errorf("err = %v, want errors.Is boom", err)
		}
	})

	t.Run("invalid attempts", func(t *testing.T) {
		err := Retry(0, func() error { return nil })
		if err == nil {
			t.Error("err = nil, want non-nil")
		}
	})
}
