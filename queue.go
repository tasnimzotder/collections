package collections

import "fmt"

type nodeQueue struct {
	value interface{}
	next  *nodeQueue
}

type Queue struct {
	front *nodeQueue
	back  *nodeQueue
	size  int
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}

// ToDo: implement enqueue and dequeue

func (q *Queue) Enqueue(value interface{}) {
	n := &nodeQueue{value, nil}

	if q.IsEmpty() {
		q.front = n
		q.back = n
	} else {
		q.back.next = n
		q.back = n
	}

	q.size += 1
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	value := q.front.value
	q.front = q.front.next
	q.size -= 1

	if q.IsEmpty() {
		q.back = nil
	}

	return value
}

func (q *Queue) Clear() {
	q.front = nil
	q.back = nil
	q.size = 0
}

func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.front.value
}

func (q *Queue) Print() {
	for n := q.front; n != nil; n = n.next {
		fmt.Println(n.value)
	}
}
