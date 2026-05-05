package syntax

import "testing"

func TestAbs(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{0, 0},
		{5, 5},
		{-5, 5},
		{-1, 1},
	}
	for _, tc := range cases {
		if got := Abs(tc.in); got != tc.want {
			t.Errorf("Abs(%d) = %d, want %d", tc.in, got, tc.want)
		}
	}
}

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{30, "FizzBuzz"},
		{7, "7"},
	}
	for _, tc := range cases {
		if got := FizzBuzz(tc.in); got != tc.want {
			t.Errorf("FizzBuzz(%d) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestSumTo(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{0, 0},
		{-3, 0},
		{1, 1},
		{5, 15},
		{10, 55},
	}
	for _, tc := range cases {
		if got := SumTo(tc.in); got != tc.want {
			t.Errorf("SumTo(%d) = %d, want %d", tc.in, got, tc.want)
		}
	}
}

func TestIsLeapYear(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{2000, true},
		{1900, false},
		{2024, true},
		{2023, false},
		{2400, true},
		{2100, false},
	}
	for _, tc := range cases {
		if got := IsLeapYear(tc.in); got != tc.want {
			t.Errorf("IsLeapYear(%d) = %v, want %v", tc.in, got, tc.want)
		}
	}
}

func TestGrade(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{100, "A"},
		{90, "A"},
		{89, "B"},
		{75, "C"},
		{61, "D"},
		{0, "F"},
		{-1, "invalid"},
		{101, "invalid"},
	}
	for _, tc := range cases {
		if got := Grade(tc.in); got != tc.want {
			t.Errorf("Grade(%d) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
