package syntax

// Abs returns the absolute value of n.
func Abs(n int) int {
	// TODO
	return 0
}

// FizzBuzz returns:
//   - "FizzBuzz" if n is divisible by both 3 and 5
//   - "Fizz" if divisible by 3 only
//   - "Buzz" if divisible by 5 only
//   - the number as a decimal string otherwise (e.g. "7")
//
// Hint: import "strconv" and use strconv.Itoa.
func FizzBuzz(n int) string {
	// TODO
	return ""
}

// SumTo returns 1 + 2 + ... + n. SumTo(0) and SumTo(negative) return 0.
// Use a for loop (don't use the closed-form n*(n+1)/2 — practice the loop).
func SumTo(n int) int {
	// TODO
	return 0
}

// IsLeapYear reports whether year is a Gregorian leap year.
// Rules: divisible by 4, EXCEPT century years not divisible by 400.
//   - 2000 → true
//   - 1900 → false
//   - 2024 → true
//   - 2023 → false
func IsLeapYear(year int) bool {
	// TODO
	return false
}

// Grade returns a letter grade for a 0-100 score.
//   - >=90: "A"
//   - >=80: "B"
//   - >=70: "C"
//   - >=60: "D"
//   - else: "F"
//
// For scores outside [0, 100], return "invalid".
func Grade(score int) string {
	// TODO
	return ""
}
