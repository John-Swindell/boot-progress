package structs

import (
	"math"
	"reflect"
	"testing"
)

func TestCounter(t *testing.T) {
	var c Counter
	if c.Value() != 0 {
		t.Errorf("zero Counter Value = %d, want 0", c.Value())
	}
	c.Inc()
	c.Inc()
	c.Add(5)
	if c.Value() != 7 {
		t.Errorf("Value = %d, want 7", c.Value())
	}
	c.Reset()
	if c.Value() != 0 {
		t.Errorf("after Reset Value = %d, want 0", c.Value())
	}
}

func TestPoint(t *testing.T) {
	p := Point{0, 0}
	q := Point{3, 4}
	if d := p.Distance(q); math.Abs(d-5) > 1e-9 {
		t.Errorf("Distance = %f, want 5", d)
	}
	r := p.Translate(1, 2)
	if r != (Point{1, 2}) {
		t.Errorf("Translate = %+v, want {1,2}", r)
	}
	if p != (Point{0, 0}) {
		t.Errorf("Translate mutated receiver: %+v", p)
	}
}

func TestStack(t *testing.T) {
	var s Stack
	if s.Len() != 0 {
		t.Errorf("empty Len = %d, want 0", s.Len())
	}
	if _, ok := s.Pop(); ok {
		t.Errorf("Pop on empty returned ok=true")
	}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Len() != 3 {
		t.Errorf("Len = %d, want 3", s.Len())
	}
	for _, want := range []int{3, 2, 1} {
		got, ok := s.Pop()
		if !ok || got != want {
			t.Errorf("Pop = (%d,%v), want (%d,true)", got, ok, want)
		}
	}
	if s.Len() != 0 {
		t.Errorf("after drain Len = %d, want 0", s.Len())
	}
}

func TestAuditCounter(t *testing.T) {
	a := NewAuditCounter()
	if a == nil || a.Counter == nil {
		t.Fatalf("NewAuditCounter returned %+v", a)
	}
	a.Inc()
	a.Inc()
	a.Inc()
	if a.Value() != 3 {
		t.Errorf("Value = %d, want 3", a.Value())
	}
	if !reflect.DeepEqual(a.Ticks, []int{1, 2, 3}) {
		t.Errorf("Ticks = %v, want [1 2 3]", a.Ticks)
	}
}
