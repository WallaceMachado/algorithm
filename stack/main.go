package main

import (
	"fmt"
	implement "github.com/WallaceMachado/algorithm/stack/implement"
)

func main() {
	s := implement.Stack{}
	fmt.Println(s.GetSize())
	fmt.Println(s.ToString())
	fmt.Println(s.Peek())
	s.Pop()
	s.Push("1")
	s.Push("2")
	s.Push("3")
	s.Push("4")
	s.Push("5")
	fmt.Println(s.GetSize())
	fmt.Println(s.ToString())
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s.GetSize())
	fmt.Println(s.ToString())
	fmt.Println(s.Peek())
}
