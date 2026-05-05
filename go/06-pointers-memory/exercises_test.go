package pointers

import (
	"reflect"
	"testing"
)

func TestIncrement(t *testing.T) {
	x := 5
	Increment(&x)
	if x != 6 {
		t.Errorf("x = %d, want 6", x)
	}
	Increment(nil) // must not panic
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	Swap(&a, &b)
	if a != 2 || b != 1 {
		t.Errorf("a,b = %d,%d, want 2,1", a, b)
	}
	Swap(nil, &a)
	Swap(&a, nil)
	Swap(nil, nil) // none of these may panic
}

func TestMaxIndex(t *testing.T) {
	cases := []struct {
		in      []int
		wantNil bool
		wantIdx int
	}{
		{nil, true, 0},
		{[]int{}, true, 0},
		{[]int{7}, false, 0},
		{[]int{1, 5, 3, 5, 2}, false, 1},
		{[]int{-3, -1, -2}, false, 1},
	}
	for _, tc := range cases {
		got := MaxIndex(tc.in)
		if tc.wantNil {
			if got != nil {
				t.Errorf("MaxIndex(%v) = %d, want nil", tc.in, *got)
			}
			continue
		}
		if got == nil {
			t.Errorf("MaxIndex(%v) = nil, want %d", tc.in, tc.wantIdx)
			continue
		}
		if *got != tc.wantIdx {
			t.Errorf("MaxIndex(%v) = %d, want %d", tc.in, *got, tc.wantIdx)
		}
	}
}

func TestLinkedList(t *testing.T) {
	if Length(nil) != 0 {
		t.Errorf("Length(nil) != 0")
	}

	h := Append(nil, 1)
	h = Append(h, 2)
	h = Append(h, 3)

	if Length(h) != 3 {
		t.Errorf("Length = %d, want 3", Length(h))
	}
	if !reflect.DeepEqual(ToSlice(h), []int{1, 2, 3}) {
		t.Errorf("ToSlice = %v, want [1 2 3]", ToSlice(h))
	}
}
