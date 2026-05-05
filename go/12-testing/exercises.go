package testlab

// Reverse returns s reversed by rune (so it handles multi-byte UTF-8 correctly).
//
//	Reverse("")        == ""
//	Reverse("a")       == "a"
//	Reverse("hello")   == "olleh"
//	Reverse("go👍")    == "👍og"   <-- 👍 is a single rune (4 UTF-8 bytes)
//
// Hint: convert to []rune, swap, convert back to string.
func Reverse(s string) string {
	// TODO
	return ""
}

// IsPalindrome reports whether s reads the same forwards and backwards by rune.
// Case-sensitive. Empty string is considered a palindrome.
func IsPalindrome(s string) bool {
	// TODO
	return false
}
