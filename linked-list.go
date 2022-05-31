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

func (l *LinkedList) Push(value interface{}) {
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

func (l *LinkedList) Pop() interface{} {
	if l.IsEmpty() {
		return nil
	}

	value := l.head.value
	l.head = l.head.next
	l.size -= 1

	if l.IsEmpty() {
		l.tail = nil
	}

	return value
}

func (l *LinkedList) Peek() interface{} {
	if l.IsEmpty() {
		return nil
	}
	return l.head.value
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
