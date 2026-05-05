package collections

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	want := []int{5, 4, 3, 2, 1}
	got := Reverse(in)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Reverse = %v, want %v", got, want)
	}
	// must not mutate input
	if !reflect.DeepEqual(in, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Reverse mutated input: %v", in)
	}

	if got := Reverse(nil); len(got) != 0 {
		t.Errorf("Reverse(nil) len = %d, want 0", len(got))
	}
}

func TestUnique(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{nil, nil},
		{[]int{}, nil},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 1, 3, 2, 4}, []int{1, 2, 3, 4}},
		{[]int{5, 5, 5}, []int{5}},
	}
	for _, tc := range cases {
		got := Unique(tc.in)
		if len(got) == 0 && len(tc.want) == 0 {
			continue
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Unique(%v) = %v, want %v", tc.in, got, tc.want)
		}
	}
}

func TestWordCount(t *testing.T) {
	got := WordCount("the quick the lazy quick fox")
	want := map[string]int{"the": 2, "quick": 2, "lazy": 1, "fox": 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WordCount = %v, want %v", got, want)
	}

	if got := WordCount(""); got == nil || len(got) != 0 {
		t.Errorf("WordCount(\"\") = %v, want non-nil empty", got)
	}
}

func TestMergeMaps(t *testing.T) {
	a := map[string]int{"x": 1, "y": 2}
	b := map[string]int{"y": 99, "z": 3}
	got := MergeMaps(a, b)
	want := map[string]int{"x": 1, "y": 99, "z": 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MergeMaps = %v, want %v", got, want)
	}
	// inputs untouched
	if a["y"] != 2 || b["y"] != 99 {
		t.Errorf("MergeMaps mutated inputs: a=%v b=%v", a, b)
	}
}

func TestTop(t *testing.T) {
	counts := map[string]int{
		"alice":   5,
		"bob":     3,
		"charlie": 5,
		"dave":    1,
	}
	cases := []struct {
		n    int
		want []string
	}{
		{0, []string{}},
		{1, []string{"alice"}},
		{2, []string{"alice", "charlie"}},
		{3, []string{"alice", "charlie", "bob"}},
		{99, []string{"alice", "charlie", "bob", "dave"}},
	}
	for _, tc := range cases {
		got := Top(counts, tc.n)
		if len(got) == 0 && len(tc.want) == 0 {
			continue
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Top(_, %d) = %v, want %v", tc.n, got, tc.want)
		}
	}
}
