package implement

import (
	"errors"
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

type LinkedList struct {
	Head *Node
	size int
}

func (l *LinkedList) GetSize() int {
	return l.size
}

// O(n) complexity, worst case scenario runs through the entire linkedlist
func (l *LinkedList) Append(val string) {
	if l.Head != nil {
		pointer := l.Head
		for pointer.Next != nil {
			pointer = pointer.Next
		}
		pointer.Next = NewNode(val)
	} else {
		l.Head = NewNode(val)

	}
	l.size = l.size + 1

}

// O(n) complexity, worst case scenario runs through the entire linkedlist
func (l *LinkedList) GetItem(index int) (string, error) {
	pointer := l.Head
	for i := 1; i <= index; i++ {
		if pointer == nil {
			return "", errors.New("index out of range")
		}
		pointer = pointer.Next
	}
	if pointer == nil {
		return "", errors.New("index out of range")
	}
	return pointer.Val, nil
}

// O(n) complexity, worst case scenario runs through the entire linkedlist
func (l *LinkedList) SetItem(index int, val string) error {
	pointer := l.Head
	for i := 1; i <= index; i++ {
		if pointer == nil {
			return errors.New("index out of range")
		}
		pointer = pointer.Next
	}
	if pointer == nil {
		return errors.New("index out of range")
	}
	pointer.Val = val
	return nil
}

// O(n) complexity, worst case scenario runs through the entire linkedlist
func (l *LinkedList) GetIndex(val string) int {
	pointer := l.Head
	i := 0
	for pointer != nil {
		if pointer.Val == val {
			return i
		}
		i++
		pointer = pointer.Next
	}
	return -1
}

// O(n) complexity, worst case scenario runs through the entire linkedlist
func (l *LinkedList) InsertItem(index int, val string) error {
	pointer := l.Head
	node := NewNode(val)
	if index == 0 {
		node.Next = l.Head
		l.Head = node
	} else {
		for i := 1; i <= index-1; i++ {
			if pointer == nil {
				return errors.New("index out of range")
			}
			pointer = pointer.Next
		}
		if pointer == nil {
			return errors.New("index out of range")
		}
		node.Next = pointer.Next
		pointer.Next = node
	}
	l.size = l.size + 1
	return nil
}

// O(n) complexity, worst case scenario runs through the entire linkedlist
func (l *LinkedList) RemoveByIndex(index int) error {
	pointer := l.Head
	if index == 0 {
		l.Head = l.Head.Next
	} else {
		for i := 1; i <= index-1; i++ {
			if pointer == nil {
				return errors.New("index out of range")
			}
			pointer = pointer.Next
		}
		if pointer == nil {
			return errors.New("index out of range")
		}
		pointer.Next = pointer.Next.Next
	}
	l.size = l.size - 1
	return nil
}

// O(n) complexity, worst case scenario runs through the entire linkedlist
func (l *LinkedList) RemoveByValue(val string) bool {
	if l.Head == nil {
		return false
	}
	if l.Head.Val == val {
		l.Head = l.Head.Next
		l.size = l.size - 1
		return true
	} else {
		ancestor := l.Head
		pointer := l.Head.Next
		for pointer != nil {
			if pointer.Val == val {
				ancestor.Next = pointer.Next
				l.size = l.size - 1
				return true
			}
			ancestor = pointer
			pointer = pointer.Next
		}
	}
	return true
}

// O(n) complexity, worst case scenario runs through the entire linkedlist
func (l *LinkedList) ToString() string {
	s := ""
	pointer := l.Head
	for pointer != nil {
		s = s + fmt.Sprintf("%v->", pointer.Val)
		pointer = pointer.Next
	}
	return s
}
