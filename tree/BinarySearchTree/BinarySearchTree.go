package tree

import "github.com/WallaceMachado/algorithm/tree/BinaryTree"

type BinarySearchTree struct {
	BinaryTree.BinaryTree
}

func (b *BinarySearchTree) Insert(value int) {
	var parent *BinaryTree.TreeNode
	x := b.Root
	// buscar a menor folha
	for x != nil {
		parent = x
		if value < x.ValNumber {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	if parent == nil {
		b.Root = BinaryTree.NewTreeNode("")
	} else if value < parent.ValNumber {
		parent.Left = BinaryTree.NewTreeNode("")
	} else {
		parent.Right = BinaryTree.NewTreeNode("")
	}
}
