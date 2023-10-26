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
	fmt.Println(s.GetSize())
	fmt.Println(s.ToString())
	fmt.Println(s.Peek())
	s.Push("2")
	fmt.Println(s.GetSize())
	fmt.Println(s.ToString())
	fmt.Println(s.Peek())
	s.Push("3")
	fmt.Println(s.GetSize())
	fmt.Println(s.ToString())
	fmt.Println(s.Peek())
	s.Push("4")
	fmt.Println(s.GetSize())
	fmt.Println(s.ToString())
	fmt.Println(s.Peek())
	s.Push("5")
	fmt.Println(s.GetSize())
	fmt.Println(s.ToString())
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s.GetSize())
	fmt.Println(s.ToString())
	fmt.Println(s.Peek())
}
