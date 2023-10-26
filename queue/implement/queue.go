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

type Queue struct {
	First *Node
	Last  *Node
	size  int
}

func (q *Queue) GetSize() int {
	return q.size
}

// O(1) complexity
func (q *Queue) ToString() string {
	str := ""
	pointer := q.First
	for pointer != nil {
		str = str + fmt.Sprintf("%v ", pointer.Val)
		pointer = pointer.Next
	}
	return str
}

// O(1) complexity
func (q *Queue) Push(val string) {
	node := NewNode(val)
	if q.First == nil {
		q.First = node

	}
	if q.Last == nil {
		q.Last = node
	} else {
		q.Last.Next = node
		q.Last = node
	}
	q.size++
}

// O(1) complexity
func (q *Queue) Pop() string {
	if q.First != nil {
		node := q.First
		q.First = node.Next
		q.size--
		return node.Val
	}
	return ""
}

// O(1) complexity
func (q *Queue) Peek() string {
	if q.size > 0 {
		return q.First.Val
	}
	return ""
}
