package tools

// LinkedListNode is a node in a doubly-linked list
type LinkedListNode struct {
	Next *LinkedListNode
	Prev *LinkedListNode
	Value int
	Weight int
}

// Insert inserts a new value to the front of the specified node and returns the new node
func (a *LinkedListNode) Insert(value int, weight int) *LinkedListNode {
	c := a.Next
	b := &LinkedListNode{
		Value: value,
		Weight: weight,
		Next: c,
		Prev: a,
	}

	a.Next = b

	if c != nil {
		c.Prev = b
	}

	return b
}

// NewLinkedList Initializes a new head node of a doubly-linked list
func NewLinkedList(root int, rootWeight int) *LinkedListNode {
	return &LinkedListNode{
		Value: root,
		Weight: rootWeight,
	}
}