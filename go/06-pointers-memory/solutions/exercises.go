package solutions

func Increment(p *int) {
	if p == nil {
		return
	}
	*p++
}

func Swap(a, b *int) {
	if a == nil || b == nil {
		return
	}
	*a, *b = *b, *a
}

func MaxIndex(xs []int) *int {
	if len(xs) == 0 {
		return nil
	}
	idx := 0
	for i, v := range xs[1:] {
		if v > xs[idx] {
			idx = i + 1
		}
	}
	return &idx
}

type Node struct {
	Val  int
	Next *Node
}

func Length(head *Node) int {
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}
	return n
}

func Append(head *Node, v int) *Node {
	node := &Node{Val: v}
	if head == nil {
		return node
	}
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
	return head
}

func ToSlice(head *Node) []int {
	var out []int
	for cur := head; cur != nil; cur = cur.Next {
		out = append(out, cur.Val)
	}
	return out
}
