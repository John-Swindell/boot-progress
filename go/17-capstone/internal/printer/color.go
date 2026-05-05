package printer

import "hash/fnv"

// ANSI 256-color foreground palette — pleasant subset, avoids dim/whitewash.
var palette = []int{
	39, 41, 43, 45, 81, 82, 83, 85, 117, 118, 119, 121,
	153, 154, 155, 157, 196, 202, 208, 214, 220, 226, 33, 51,
}

// colorFor returns a stable ANSI 256-color escape sequence for the given key.
// Hash → palette index → "\x1b[38;5;Nm".
func colorFor(key string) string {
	h := fnv.New32a()
	_, _ = h.Write([]byte(key))
	c := palette[int(h.Sum32())%len(palette)]
	return ansi256(c)
}

func ansi256(n int) string {
	return "\x1b[38;5;" + itoa(n) + "m"
}

const reset = "\x1b[0m"

// itoa: tiny no-alloc int->string for [0,999].
func itoa(n int) string {
	if n < 10 {
		return string(rune('0' + n))
	}
	if n < 100 {
		return string([]byte{byte('0' + n/10), byte('0' + n%10)})
	}
	return string([]byte{byte('0' + n/100), byte('0' + (n/10)%10), byte('0' + n%10)})
}
