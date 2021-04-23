package practice

import (
	"fmt"
	"testing"
)

func TestInorderRecursive(t *testing.T) {
	// InorderTest inorder is from buttom, left node => parent node => right node
	/*
				10
			   /  \
			 20	   30
			/ \      \
		   40  50     60
		  /
		 70
	*/

	root := &Node{10, nil, nil}
	root.left = &Node{20, nil, nil}
	root.right = &Node{30, nil, nil}
	root.left.left = &Node{40, nil, nil}
	root.left.right = &Node{50, nil, nil}
	root.right.right = &Node{60, nil, nil}
	root.left.left.left = &Node{70, nil, nil}

	fmt.Println("Inorder Traversal - recursive solution : ")
	InorderRecursive(root)
}

func TestPostorderRecursive(t *testing.T) {
	// Postorder is buttom first left-right from buttom start from left, left => right => parent
	/*
				10
			   /  \
			 20	   30
			/ \      \
		   40  50     60
		  /
		 70
	*/

	root := &Node{10, nil, nil}
	root.left = &Node{20, nil, nil}
	root.right = &Node{30, nil, nil}
	root.left.left = &Node{40, nil, nil}
	root.left.right = &Node{50, nil, nil}
	root.right.right = &Node{60, nil, nil}
	root.left.left.left = &Node{70, nil, nil}

	fmt.Println("Postorder Traversal - recursive solution : ")
	PostorderRecursive(root)
}
func TestPreorderRecursive(t *testing.T) {
	// Preorder is from top to buttom start left, parent => left (til end) => back to unvisited right node
	/*
				10
			   /  \
			 20	   30
			/ \      \
		   40  50     60
		  /
		 70
	*/

	root := &Node{10, nil, nil}
	root.left = &Node{20, nil, nil}
	root.right = &Node{30, nil, nil}
	root.left.left = &Node{40, nil, nil}
	root.left.right = &Node{50, nil, nil}
	root.right.right = &Node{60, nil, nil}
	root.left.left.left = &Node{70, nil, nil}

	fmt.Println("Preorder Traversal - recursive solution : ")
	PreorderRecursive(root)
}
