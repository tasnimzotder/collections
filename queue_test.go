package collections

import "testing"

func TestQueue(t *testing.T) {
	var q Queue

	if q.Size() != 0 {
		t.Errorf("Length of queue should be 0, but got %d", q.Size())
	}

	q.Push(1)

	if q.Size() != 1 {
		t.Errorf("Length of queue should be 1, but got %d", q.Size())
	}

	if q.Peek() != 1 {
		t.Errorf("Peek should return 1, but got %d", q.Peek())
	}

	q.Push(2)

	if q.Size() != 2 {
		t.Errorf("Length of queue should be 2, but got %d", q.Size())
	}

	if q.Peek() != 1 {
		t.Errorf("Peek should return 1, but got %d", q.Peek())
	}

	if q.Pop() != 1 {
		t.Errorf("Dequeue should return 1, but got %d", q.Pop())
	}

	if q.Size() != 1 {
		t.Errorf("Length of queue should be 1, but got %d", q.Size())
	}

	if q.Peek() != 2 {
		t.Errorf("Peek should return 2, but got %d", q.Peek())
	}

	if q.Pop() != 2 {
		t.Errorf("Dequeue should return 2, but got %d", q.Pop())
	}

	if q.Size() != 0 {
		t.Errorf("Length of queue should be 0, but got %d", q.Size())
	}

	if !q.IsEmpty() {
		t.Errorf("Queue should be empty, but got %d", q.Size())
	}
}
