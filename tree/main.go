package main

import (
	"fmt"
	tree "github.com/WallaceMachado/algorithm/tree/BinaryTree"
)

//    '+'
//   /   \
// 'a'   '*'
//   	/   \
// 	  'b'   '-'
//   	   /   \
// 		 '/'   'e'
//   	/   \
// 	  'c'   'd'
// (a + (b * ((c/d) - e)))

func main() {
	a := tree.NewTreeNode("a")
	plus := tree.NewTreeNode("+")
	b := tree.NewTreeNode("b")
	times := tree.NewTreeNode("*")
	sub := tree.NewTreeNode("-")
	div := tree.NewTreeNode("/")
	c := tree.NewTreeNode("c")
	d := tree.NewTreeNode("d")
	e := tree.NewTreeNode("e")
	div.Left = c
	div.Right = d
	sub.Left = div
	sub.Right = e
	times.Left = b
	times.Right = sub
	plus.Left = a
	plus.Right = times
	tree := tree.NewBinaryTree(plus)
	tree.SimetricTraversal(nil)
	fmt.Println()
	tree.PosOrderTraversal(nil)
	fmt.Println()
	fmt.Println(tree.Height(nil))
	fmt.Println(tree.Height(times))
	fmt.Println(tree.Height(sub))
	fmt.Println(tree.Height(a))
}
