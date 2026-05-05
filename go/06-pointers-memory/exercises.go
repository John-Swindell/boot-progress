package pointers

// Increment adds 1 to the int that p points to.
// If p is nil, do nothing.
func Increment(p *int) {
	// TODO
}

// Swap exchanges the values pointed to by a and b. If either is nil, do nothing.
func Swap(a, b *int) {
	// TODO
}

// MaxIndex returns a pointer to the index of the largest element in xs.
// On ties, the smallest index wins.
// On empty input, returns nil.
//
// Hint: declare an int, take its address, return that.
func MaxIndex(xs []int) *int {
	// TODO
	return nil
}

// Node is a singly-linked list node.
type Node struct {
	Val  int
	Next *Node
}

// Length returns the number of nodes in the list starting at head.
// Length(nil) == 0.
func Length(head *Node) int {
	// TODO
	return 0
}

// Append adds a new node with value v to the END of the list. It returns
// the (possibly-new) head.
//
// Edge case: Append(nil, 5) returns a one-node list whose head Val == 5.
func Append(head *Node, v int) *Node {
	// TODO
	return nil
}

// ToSlice flattens the linked list into a []int in order.
func ToSlice(head *Node) []int {
	// TODO
	return nil
}
