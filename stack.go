package collections

import "fmt"

type nodeStack struct {
	value interface{}
	next  *nodeStack
}

type Stack struct {
	top  *nodeStack
	size int
}

func (s *Stack) Push(value interface{}) {
	s.top = &nodeStack{value, s.top}
	s.size += 1
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	value := s.top.value
	s.top = s.top.next
	s.size -= 1
	return value
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.top.value
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Clear() {
	s.top = nil
	s.size = 0
}

func (s *Stack) Print() {
	for n := s.top; n != nil; n = n.next {
		fmt.Println(n.value)
	}
}
