package algorithm

import (
	"fmt"
	"testing"
)

func TestBFSDFS(t *testing.T) {
	numbers := []int{53, 23, 19, 5776, 170, 223, 45, 75, 90, 802, 63, 29, 3, 887, 456, 24, 2, 21, 34, 49, 6555}
	tree := BinaryTree(numbers)
	fmt.Println(tree.BFS())
	// fmt.Println(tree.Find(49))
	// fmt.Println(tree.Maximum())
	// fmt.Println(tree.Minimum())
	// fmt.Println("del", tree.Delete(223))
	// fmt.Println(tree.BFS())

	fmt.Println(tree.DFS())
}
