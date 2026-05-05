package goroutines

import (
	"sort"
	"sync"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	got := []int{}
	for v := range Generate(1, 2, 3, 4) {
		got = append(got, v)
	}
	want := []int{1, 2, 3, 4}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func TestSquare(t *testing.T) {
	out := Square(Generate(1, 2, 3, 4))
	got := []int{}
	for v := range out {
		got = append(got, v)
	}
	want := []int{1, 4, 9, 16}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func TestSum(t *testing.T) {
	if got := Sum(Generate(1, 2, 3, 4, 5)); got != 15 {
		t.Errorf("Sum = %d, want 15", got)
	}
}

func TestMerge(t *testing.T) {
	a := Generate(1, 2, 3)
	b := Generate(10, 20)
	c := Generate(100)
	out := Merge(a, b, c)

	got := []int{}
	for v := range out {
		got = append(got, v)
	}
	sort.Ints(got)
	want := []int{1, 2, 3, 10, 20, 100}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func TestFirstResponse(t *testing.T) {
	slow := make(chan int)
	fast := make(chan int)
	go func() {
		time.Sleep(50 * time.Millisecond)
		slow <- 999
		close(slow)
	}()
	go func() {
		fast <- 42
		close(fast)
	}()

	v, ok := FirstResponse([]<-chan int{slow, fast})
	if !ok || v != 42 {
		t.Errorf("FirstResponse = (%d,%v), want (42,true)", v, ok)
	}

	// All-closed-empty case
	empty1 := make(chan int)
	empty2 := make(chan int)
	close(empty1)
	close(empty2)
	v, ok = FirstResponse([]<-chan int{empty1, empty2})
	if ok {
		t.Errorf("expected (0,false), got (%d,true)", v)
	}
}

func TestWithTimeout(t *testing.T) {
	t.Run("times out", func(t *testing.T) {
		ch := make(chan int)
		v, ok := WithTimeout(ch, 10*time.Millisecond)
		if ok {
			t.Errorf("expected timeout, got (%d,true)", v)
		}
	})
	t.Run("receives in time", func(t *testing.T) {
		ch := make(chan int, 1)
		ch <- 7
		v, ok := WithTimeout(ch, 50*time.Millisecond)
		if !ok || v != 7 {
			t.Errorf("got (%d,%v), want (7,true)", v, ok)
		}
	})
}

// keep the linter happy if we don't end up using it elsewhere
var _ = sync.WaitGroup{}
