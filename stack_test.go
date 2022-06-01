package collections

import "testing"

func TestStack(t *testing.T) {
	var s Stack

	if s.Size() != 0 {
		t.Errorf("Length of stack should be 0, but got %d", s.Size())
	}

	s.Push(1)

	if s.Size() != 1 {
		t.Errorf("Length of stack should be 1, but got %d", s.Size())
	}

	if s.Peek() != 1 {
		t.Errorf("Peek should return 1, but got %d", s.Peek())
	}

	s.Push(2)

	if s.Size() != 2 {
		t.Errorf("Length of stack should be 2, but got %d", s.Size())
	}

	if s.Peek() != 2 {
		t.Errorf("Peek should return 2, but got %d", s.Peek())
	}

	if s.Pop() != 2 {
		t.Errorf("Pop should return 2, but got %d", s.Pop())
	}

	if s.Size() != 1 {
		t.Errorf("Length of stack should be 1, but got %d", s.Size())
	}

	if s.Peek() != 1 {
		t.Errorf("Peek should return 1, but got %d", s.Peek())
	}

	if s.Pop() != 1 {
		t.Errorf("Pop should return 1, but got %d", s.Pop())
	}

	if s.Size() != 0 {
		t.Errorf("Length of stack should be 0, but got %d", s.Size())
	}

	if !s.IsEmpty() {
		t.Errorf("Stack should be empty, but got %d", s.Size())
	}
}
