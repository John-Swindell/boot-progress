package testlab

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		name, in, want string
	}{
		{"empty", "", ""},
		{"single", "a", "a"},
		{"ascii", "hello", "olleh"},
		{"emoji", "go👍", "👍og"},
		{"multibyte", "héllo", "olléh"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Reverse(tc.in); got != tc.want {
				t.Errorf("Reverse(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"aba", true},
		{"abc", false},
		{"Aba", false}, // case sensitive
		{"racecar", true},
	}
	for _, tc := range cases {
		if got := IsPalindrome(tc.in); got != tc.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", tc.in, got, tc.want)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	s := "the quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Reverse(s)
	}
}

func FuzzReverseRoundtrip(f *testing.F) {
	f.Add("hello")
	f.Add("")
	f.Add("👍")
	f.Add("héllo")
	f.Fuzz(func(t *testing.T, s string) {
		if got := Reverse(Reverse(s)); got != s {
			t.Errorf("roundtrip failed for %q -> %q", s, got)
		}
	})
}
