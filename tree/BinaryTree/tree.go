package BinaryTree

import "fmt"

// na arvore binaria de busca os valores da esquerda sção sempre menores que o no e o das direta são maiores

type TreeNode struct {
	Val       string
	ValNumber int
	Right     *TreeNode
	Left      *TreeNode
}

func NewTreeNode(val string) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

type BinaryTree struct {
	Root *TreeNode
}

func NewBinaryTree(n *TreeNode) *BinaryTree {
	return &BinaryTree{Root: n}
}

func (t *TreeNode) ToString() string {
	return t.Val
}

var i int

// percurso em ordem simétrica (o correto em InOrder em ingles)
func (t *BinaryTree) SimetricTraversal(node *TreeNode) {
	if node == nil {
		node = t.Root
	}
	if node.Left != nil {
		// parenteses são específicos para o nosso exemplo
		// um percurso em ordem simétrica não precisa deles
		fmt.Print("(")
		t.SimetricTraversal(node.Left)
	}
	fmt.Print(node.ToString())
	if node.Right != nil {

		t.SimetricTraversal(node.Right)
		fmt.Print(")")
	}
}

func (t *BinaryTree) PosOrderTraversal(node *TreeNode) {
	if node == nil {
		node = t.Root
	}
	if node.Left != nil {
		t.PosOrderTraversal(node.Left)
	}
	if node.Right != nil {
		t.PosOrderTraversal(node.Right)
	}
	fmt.Print(node.ToString())
}

func (t *BinaryTree) Height(node *TreeNode) int {
	if node == nil {
		node = t.Root
	}
	hLeft := 0
	hRight := 0
	if node.Left != nil {
		hLeft = t.Height(node.Left)
	}
	if node.Right != nil {
		hRight = t.Height(node.Right)
	}
	if hRight > hLeft {
		return hRight + 1
	}
	return hLeft + 1
}
