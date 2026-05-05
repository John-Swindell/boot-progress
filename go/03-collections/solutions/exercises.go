package solutions

import (
	"sort"
	"strings"
)

func Reverse(xs []int) []int {
	out := make([]int, len(xs))
	for i, v := range xs {
		out[len(xs)-1-i] = v
	}
	return out
}

func Unique(xs []int) []int {
	if len(xs) == 0 {
		return nil
	}
	seen := make(map[int]struct{}, len(xs))
	out := make([]int, 0, len(xs))
	for _, x := range xs {
		if _, ok := seen[x]; ok {
			continue
		}
		seen[x] = struct{}{}
		out = append(out, x)
	}
	return out
}

func WordCount(s string) map[string]int {
	out := make(map[string]int)
	for _, w := range strings.Fields(s) {
		out[w]++
	}
	return out
}

func MergeMaps(a, b map[string]int) map[string]int {
	out := make(map[string]int, len(a)+len(b))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		out[k] = v
	}
	return out
}

func Top(counts map[string]int, n int) []string {
	if n <= 0 {
		return nil
	}
	type kv struct {
		k string
		v int
	}
	items := make([]kv, 0, len(counts))
	for k, v := range counts {
		items = append(items, kv{k, v})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].v != items[j].v {
			return items[i].v > items[j].v
		}
		return items[i].k < items[j].k
	})
	if n > len(items) {
		n = len(items)
	}
	out := make([]string, n)
	for i := 0; i < n; i++ {
		out[i] = items[i].k
	}
	return out
}
