package main

import (
	"fmt"
	"github.com/WallaceMachado/algorithm/queue/implement"
)

func main() {
	q := implement.Queue{}
	fmt.Println("--------------------")
	fmt.Println(q.GetSize())
	fmt.Println(q.ToString())
	fmt.Println(q.Peek())
	q.Pop()
	fmt.Println("--------------------")
	q.Push("1")
	q.Push("2")
	q.Push("3")
	q.Push("4")
	q.Push("5")
	fmt.Println("--------------------")
	fmt.Println(q.GetSize())
	fmt.Println(q.ToString())
	fmt.Println(q.Peek())
	fmt.Println("--------------------")
	fmt.Println(q.Pop())
	fmt.Println(q.GetSize())
	fmt.Println(q.ToString())
	fmt.Println(q.Peek())
	fmt.Println("--------------------")
}
