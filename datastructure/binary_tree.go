package datastructure

// BinaryTreeNode to represent node of binary tree
type BinaryTreeNode struct {
	Key   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// Insert will add node of the tree, key should be unique
func (b *BinaryTreeNode) Insert(k int) {
	if b.Key < k {
		// move right
		if b.Right == nil {
			b.Right = &BinaryTreeNode{Key: k}
		} else {
			b.Right.Insert(k)
		}
	} else if b.Key > k {
		//move left
		if b.Left == nil {
			b.Left = &BinaryTreeNode{Key: k}
		} else {
			b.Left.Insert(k)
		}
	}

}

// Search to search value in binary tree
func (b *BinaryTreeNode) Search(k int) bool {
	if b == nil {
		return false
	}
	if b.Key < k {
		// move right
		return b.Right.Search(k)
	} else if b.Key > k {
		//move left
		return b.Left.Search(k)
	}
	// equal means found
	return true
}
