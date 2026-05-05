package structs

// ----- Counter -----

// Counter is a simple counter that mutates in place.
//
// Pick the right receiver kind for each method.
type Counter struct {
	n int
}

// Inc increments the counter by 1.
func (c *Counter) Inc() {
	// TODO
}

// Add increments the counter by delta.
func (c *Counter) Add(delta int) {
	// TODO
}

// Value returns the current count without modifying it.
func (c *Counter) Value() int {
	// TODO
	return 0
}

// Reset sets the counter back to zero.
func (c *Counter) Reset() {
	// TODO
}

// ----- Point -----

// Point is an immutable 2D point. Use value receivers for its methods.
type Point struct {
	X, Y float64
}

// Distance returns the Euclidean distance from p to other.
//
// Hint: math.Hypot(dx, dy).
func (p Point) Distance(other Point) float64 {
	// TODO
	return 0
}

// Translate returns a NEW Point shifted by (dx, dy). p is not mutated.
func (p Point) Translate(dx, dy float64) Point {
	// TODO
	return Point{}
}

// ----- Stack (generic-free, []int) -----

// Stack is a LIFO stack of ints backed by a slice.
type Stack struct {
	data []int
}

// Push appends x to the top of the stack.
func (s *Stack) Push(x int) {
	// TODO
}

// Pop removes and returns the top element.
// Returns (0, false) when the stack is empty.
func (s *Stack) Pop() (int, bool) {
	// TODO
	return 0, false
}

// Len returns the number of elements on the stack.
func (s *Stack) Len() int {
	// TODO
	return 0
}

// ----- Embedding -----

// AuditCounter embeds *Counter and additionally records every Inc call's
// "tick" — a monotonically incrementing int starting at 1.
type AuditCounter struct {
	*Counter
	Ticks []int
}

// NewAuditCounter returns an AuditCounter with a fresh embedded Counter and
// empty (non-nil) Ticks slice.
func NewAuditCounter() *AuditCounter {
	// TODO
	return nil
}

// Inc overrides the embedded Counter.Inc to ALSO append a tick to Ticks.
// The first Inc records 1, the second 2, and so on.
//
// Tip: call a.Counter.Inc() to invoke the embedded method.
func (a *AuditCounter) Inc() {
	// TODO
}
