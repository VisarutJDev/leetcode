package practice

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func InorderRecursive(root *Node) {
	if root == nil {
		return
	}

	InorderRecursive(root.left)
	fmt.Printf("%d ", root.val)
	InorderRecursive(root.right)
}

func PostorderRecursive(root *Node) {
	if root == nil {
		return
	}

	PostorderRecursive(root.left)
	PostorderRecursive(root.right)
	fmt.Printf("%d ", root.val)
}

func PreorderRecursive(root *Node) {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.val)
	PreorderRecursive(root.left)
	PreorderRecursive(root.right)
}
