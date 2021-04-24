package practice

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	n := []int{1, 2, 3, 4, 5, 6}
	// root := CreateBST(n)
	// InorderRecursive(root)
	// fmt.Println()
	// PostorderRecursive(root)
	// fmt.Println()
	// PreorderRecursive(root)
	// fmt.Println()
	root2 := CreateBSTV2(n, 0, len(n)-1)
	// InorderRecursive(root2)
	// fmt.Println()
	// PostorderRecursive(root2)
	// fmt.Println()
	// PreorderRecursive(root2)
	// fmt.Println()

	CreateLevelLinkedList(root2)
	InorderRecursive(root2)
	fmt.Println()
}
