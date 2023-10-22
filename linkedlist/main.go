package main

import (
	"fmt"
	"github.com/WallaceMachado/algorithm/linkedlist/implement"
)

func main() {
	l := implement.LinkedList{}
	fmt.Println(l)
	l.Append("1")
	fmt.Println(l)
	l.Append("2")
	fmt.Println(l)
	l.Append("3")
	fmt.Println(l)
	l.Append("4")
	fmt.Println(l)
	l.Append("5")
	fmt.Println(l)
	for i := 0; i < l.GetSize(); i++ {
		v, err := l.GetItem(i)
		fmt.Println(v, err)
	}
	fmt.Println(l.GetSize())
	err := l.InsertItem(0, "0")
	fmt.Println(err)
	v, err := l.GetItem(0)
	fmt.Println(v, err)
	for i := 0; i < l.GetSize(); i++ {
		v, err := l.GetItem(i)
		fmt.Println(v, err)
	}
	err = l.InsertItem(3, "33")
	fmt.Println(err)
	v, err = l.GetItem(3)
	fmt.Println(v, err)
	err = l.InsertItem(l.GetSize(), "6")
	fmt.Println(err)
	v, err = l.GetItem(l.GetSize() - 1)
	fmt.Println(v, err)
	fmt.Println(l.ToString())
	l.RemoveByValue("0")
	fmt.Println(l.ToString())
	l.RemoveByValue("33")
	fmt.Println(l.ToString())
	l.RemoveByValue("6")
	fmt.Println(l.ToString())
	l.RemoveByIndex(0)
	fmt.Println(l.ToString())
	l.RemoveByIndex(2)
	fmt.Println(l.ToString())
	l.RemoveByIndex(l.GetSize() - 1)
	fmt.Println(l.ToString())
}
