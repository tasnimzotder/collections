package collections

import "testing"

func TestLinkedList(t *testing.T) {
	var ll LinkedList

	if ll.Size() != 0 {
		t.Errorf("Length of linked list should be 0, but got %d", ll.Size())
	}

	ll.Push(1)

	if ll.Size() != 1 {
		t.Errorf("Length of linked list should be 1, but got %d", ll.Size())
	}

	if ll.Get(0) != 1 {
		t.Errorf("At should return 1, but got %d", ll.Get(0))
	}

	ll.Push(2)

	if ll.Size() != 2 {
		t.Errorf("Length of linked list should be 2, but got %d", ll.Size())
	}

	if ll.Pop() != 2 {
		t.Errorf("Pop should return 2, but got %d", ll.Pop())
	}

	ll.InsertAt(0, 3)

	if ll.Size() != 2 {
		t.Errorf("Length of linked list should be 2, but got %d", ll.Size())
	}

	if ll.Get(0) != 3 {
		t.Errorf("At should return 3, but got %d", ll.Get(0))
	}

	// clear
	ll.Clear()

	if ll.Size() != 0 {
		t.Errorf("Length of linked list should be 0, but got %d", ll.Size())
	}
}
