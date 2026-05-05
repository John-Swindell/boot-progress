package solutions

import "math"

type Counter struct{ n int }

func (c *Counter) Inc()          { c.n++ }
func (c *Counter) Add(delta int) { c.n += delta }
func (c *Counter) Value() int    { return c.n }
func (c *Counter) Reset()        { c.n = 0 }

type Point struct{ X, Y float64 }

func (p Point) Distance(o Point) float64       { return math.Hypot(p.X-o.X, p.Y-o.Y) }
func (p Point) Translate(dx, dy float64) Point { return Point{p.X + dx, p.Y + dy} }

type Stack struct{ data []int }

func (s *Stack) Push(x int) { s.data = append(s.data, x) }
func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	x := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return x, true
}
func (s *Stack) Len() int { return len(s.data) }

type AuditCounter struct {
	*Counter
	Ticks []int
}

func NewAuditCounter() *AuditCounter {
	return &AuditCounter{Counter: &Counter{}, Ticks: []int{}}
}
func (a *AuditCounter) Inc() {
	a.Counter.Inc()
	a.Ticks = append(a.Ticks, a.Counter.Value())
}
