package collections

import "fmt"

type nodeLinkedList struct {
	value interface{}
	next  *nodeLinkedList
}

type LinkedList struct {
	head *nodeLinkedList
	tail *nodeLinkedList
	size int
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) InsertAtEnd(value interface{}) {
	n := &nodeLinkedList{value, nil}

	if l.IsEmpty() {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}

	l.size += 1
}

func (l *LinkedList) Push(value interface{}) {
	l.InsertAtEnd(value)
}

func (l *LinkedList) InsertAtBeginning(value interface{}) {
	n := &nodeLinkedList{value, nil}

	if l.IsEmpty() {
		l.head = n
		l.tail = n
	} else {
		n.next = l.head
		l.head = n
	}

	l.size += 1
}

func (l *LinkedList) InsertAt(index int, value interface{}) {
	if index < 0 || index > l.size {
		return
	}

	if index == 0 {
		l.InsertAtBeginning(value)
		return
	}

	if index == l.size {
		l.InsertAtEnd(value)
		return
	}

	n := &nodeLinkedList{value, nil}
	current := l.head

	for i := 0; i < index-1; i++ {
		current = current.next
	}

	n.next = current.next
	current.next = n

	l.size += 1
}

func (l *LinkedList) RemoveFrom(index int) interface{} {
	if index < 0 || index >= l.size {
		return nil
	}

	if index == 0 {
		value := l.head.value
		l.head = l.head.next
		l.size -= 1

		return value
	}

	if index == l.size-1 {
		value := l.tail.value
		current := l.head

		for i := 0; i < index-1; i++ {
			current = current.next
		}

		l.tail = current
		l.tail.next = nil
		l.size -= 1
		return value
	}

	current := l.head

	for i := 0; i < index-1; i++ {
		current = current.next
	}

	value := current.next.value
	current.next = current.next.next

	l.size -= 1

	return value
}

func (l *LinkedList) Pop() interface{} {
	return l.RemoveFrom(l.size - 1)
}

func (l *LinkedList) Get(index int) interface{} {
	if index < 0 || index >= l.size {
		return nil
	}

	current := l.head

	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.value
}

func (l *LinkedList) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *LinkedList) Print() {
	for n := l.head; n != nil; n = n.next {
		fmt.Println(n.value)
	}
}
