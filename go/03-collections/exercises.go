package collections

// Reverse returns a NEW slice with the elements of xs in reverse order.
// It must NOT mutate xs.
func Reverse(xs []int) []int {
	// TODO
	return nil
}

// Unique returns a slice of the unique elements of xs in first-seen order.
// Example: Unique([]int{1, 2, 1, 3, 2, 4}) -> [1, 2, 3, 4]
func Unique(xs []int) []int {
	// TODO — use a map[int]struct{} as a set.
	return nil
}

// WordCount splits s on whitespace (use strings.Fields) and returns a map
// from word -> occurrence count.
//
// WordCount("") -> empty (non-nil) map.
func WordCount(s string) map[string]int {
	// TODO
	return nil
}

// MergeMaps returns a new map containing all keys from a and b. On key
// conflict, b's value wins. Neither input map is mutated.
//
// nil inputs are treated as empty.
func MergeMaps(a, b map[string]int) map[string]int {
	// TODO
	return nil
}

// Top returns the top-n keys of counts ordered by:
//  1. count descending
//  2. then key ascending (alphabetical) on ties
//
// If counts has fewer than n entries, returns all of them.
// n <= 0 returns an empty slice.
//
// Hint: import "sort"; build a slice of structs, sort, take first n.
func Top(counts map[string]int, n int) []string {
	// TODO
	return nil
}
