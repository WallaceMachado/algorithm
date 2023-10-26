package implement

import (
	"fmt"
)

type Node struct {
	Val  string
	Next *Node
}

func NewNode(val string) *Node {
	return &Node{
		Val: val,
	}
}

type Stack struct {
	Top  *Node
	size int
}

func (s *Stack) GetSize() int {
	return s.size
}

// O(1) complexity
func (s *Stack) ToString() string {
	str := ""
	pointer := s.Top
	for pointer != nil {
		str = str + fmt.Sprintf("%v\n", pointer.Val)
		pointer = pointer.Next
	}
	return str
}

// O(1) complexity
func (s *Stack) Push(val string) {
	node := NewNode(val)
	node.Next = s.Top
	s.Top = node
	s.size++
}

// O(1) complexity
func (s *Stack) Pop() string {
	if s.size > 0 {
		node := s.Top
		s.Top = s.Top.Next
		s.size--
		return node.Val
	}
	return ""
}

// O(1) complexity
func (s *Stack) Peek() string {
	if s.size > 0 {
		return s.Top.Val
	}
	return ""
}
