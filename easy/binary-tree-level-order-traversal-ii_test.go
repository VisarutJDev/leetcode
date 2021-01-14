package easy

import (
	"fmt"
	"testing"
)

func newLinkedList(n int) *TreeNode { return &TreeNode{Val: n} }

func TestLevelOrderBottom(t *testing.T) {
	head := newLinkedList(3)
	head.Left = newLinkedList(9)
	head.Right = newLinkedList(20)
	head.Right.Left = newLinkedList(15)
	head.Right.Right = newLinkedList(7)

	fmt.Println(LevelOrderBottom(head))
}
