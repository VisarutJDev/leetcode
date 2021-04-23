package algorithm

import (
	"errors"
	"math"
)

// READ HERE https://medium.com/@houzier.saurav/dfs-and-bfs-golang-d5818ec690d3

//Tree to print the tree in breadth first fashion
type Tree struct {
	Value  int
	Left   *Tree
	Right  *Tree
	Repeat int
	Parent *Tree
}

// BFS is breadth first search function
func (b *Tree) BFS() []int {
	queue := []*Tree{}
	queue = append(queue, b)
	result := []int{}
	return BFSUtil(queue, result)
}

//BFSUtil, This appends the value of the root of subtree or tree to the
//result
func BFSUtil(queue []*Tree, res []int) []int {
	if len(queue) == 0 {
		return res
	}
	res = append(res, queue[0].Value)
	if queue[0].Left != nil {
		queue = append(queue, queue[0].Left)
	}
	if queue[0].Right != nil {
		queue = append(queue, queue[0].Right)
	}

	return BFSUtil(queue[1:], res)
}

//BinaryTree struct which actually makes a binary tree out of the
//integer array called as numbers
func BinaryTree(numbers []int) *Tree {
	root := numbers[0]
	tree := Tree{root, nil, nil, 0, nil}
	for _, value := range numbers[1:] {
		tree.Insert(value)
	}
	return &tree
}

//Insert method to insert new values in the trees
//Sice every node is itself a Tree struct, we can use
//Tree struct of each particular node
func (t *Tree) Insert(v int) error {

	switch {

	case v == t.Value:
		t.Repeat++

	//if value is less than the value at the node and

	//if Left is empty, add a node here only.
	//if left is not empty, recrusion
	case v < t.Value:
		if t.Left == nil {
			t.Left = &Tree{v, nil, nil, 1, t}
			return nil
		}

		return t.Left.Insert(v)

	//if value is greater than the value at the node and

	//if right is empty, add a node here only.
	//if right is not empty, recrusion
	case v > t.Value:
		if t.Right == nil {
			t.Right = &Tree{v, nil, nil, 1, t}
			return nil
		}

		return t.Right.Insert(v)

	}

	return nil

}

//Find To find whether a value exists in the binary tree or
//not by scanning the tree, returns (int, bool) int represents
//number of times the value is present in the tree,
//The *Tree in the return statement is the sub tree with
//its root node as Value
func (t *Tree) Find(v int) (int, *Tree, bool) {
	switch {
	case v == t.Value:
		return t.Repeat, t, true

	case v < t.Value:

		if t.Left == nil {
			return 0, nil, false
		}

		return t.Left.Find(v)

	case v > t.Value:

		if t.Right == nil {
			return 0, nil, false
		}

		return t.Right.Find(v)

	}

	return 0, nil, false

}

//Maximum Maximum integer present in the tree,
func (t *Tree) Maximum() int {

	if t.Right == nil {
		return t.Value
	}
	return t.Right.Maximum()

}

//Minimum Minimum integer present in the tree
//If we have to ding minimum we should span the
//left subtree
func (t *Tree) Minimum() int {
	if t.Left == nil {
		return t.Value
	}
	return t.Left.Minimum()
}

//Delete  delete a value from the binary tree
func (t *Tree) Delete(v int) error {
	_, tree, present := t.Find(v)

	if !present {

		return errors.New("Element not present in the tree")
	}

	return tree.Replace(v)

}

//Replace
func (t *Tree) Replace(v int) error {

	if t.Right == nil && t.Left == nil {
		t.Parent = nil
		return nil
	}

	switch {
	case t.Right == nil:
		t.Parent.Left = t.Left

	case t.Left == nil:
		t.Parent.Right = t.Right
	default:
		//When the subtree has a value which is to be deleted but has both left and right subtree.
		//We will find the minimum in the right subtree and replace it with the value to be
		//deleted. Now we will delete this minimum in the right subtreee
		min := t.Right.Minimum()
		t.Value = min
		t.Right.Delete(min)

	}

	return nil
}

func (t *Tree) IsBalance() bool {
	return !(checkHeight(t) == -1)
}

func checkHeight(root *Tree) int {
	if root == nil {
		return 0
	}
	leftHeight := checkHeight(root.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := checkHeight(root.Right)
	if rightHeight == -1 {
		return -1
	}

	diff := leftHeight - rightHeight
	if math.Abs(float64(diff)) > 1 {
		return -1
	} else {
		return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
	}
}
