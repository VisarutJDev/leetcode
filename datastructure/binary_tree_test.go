package datastructure

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := &BinaryTreeNode{Key: 100}
	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(88)
	tree.Insert(276)

	fmt.Println(tree)
	fmt.Println(tree.Search(76))
	fmt.Println(tree.Search(1))
	fmt.Println(tree.Search(400))
}
