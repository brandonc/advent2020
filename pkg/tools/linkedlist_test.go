package tools

import "testing"

func TestLinkedList(t *testing.T) {
	head := NewLinkedList(0, 0)
	head.Insert(3, 0).Insert(6, 0).Insert(9, 0)

	if head.Value != 0 {
		t.Errorf("head == %d; wanted 0", head.Value)
	}

	if head.Next.Value != 3 {
		t.Errorf("head.Next == %d; wanted 3", head.Next.Value)
	}

	if head.Next.Next.Value != 6 {
		t.Errorf("head.Next.Next == %d; wanted 6", head.Next.Next.Value)
	}

	if head.Next.Next.Next.Value != 9 {
		t.Errorf("head.Next.Next.Next == %d; wanted 9", head.Next.Next.Next.Value)
	}

	tail := head.Next.Next.Next

	if tail.Next != nil {
		t.Errorf("tail.Next == %p; wanted nil", tail.Next)
	}

	if tail.Prev.Value != 6 {
		t.Errorf("tail.Prev == %d; wanted 6", tail.Prev.Value)
	}

	if tail.Prev.Prev.Value != 3 {
		t.Errorf("tail.Prev == %d; wanted 3", tail.Prev.Prev.Value)
	}

	if tail.Prev.Prev.Prev != head {
		t.Errorf("tail.Prev.Prev.Prev == %p; wanted %p", tail.Prev.Prev.Prev, head)
	}
}