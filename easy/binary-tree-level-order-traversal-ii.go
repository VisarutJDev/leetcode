package easy

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LevelOrderBottom question
// Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).
// Logic is to collect value in each level by queue and reverse array
func LevelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	levelLen := len(q)
	var r [][]int
	var level []int
	for len(q) != 0 {
		if levelLen == 0 {
			r = append(r, level)
			levelLen = len(q)
			level = []int{}
		}
		n := q[0]
		levelLen--
		level = append(level, n.Val)
		q = q[1:]
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
	}
	rr := [][]int{level}
	for i := len(r) - 1; i >= 0; i-- {
		rr = append(rr, r[i])
	}
	return rr
}

func walker(t *TreeNode) {
	// initiate walking over tree t with corresponding channel
	walk(t)
}

var array []*int

func walk(t *TreeNode) {
	if t == nil {
		// array = append(array, nil)
		return
	}

	// Preorder (Root, Left, Right)
	// fmt.Println("val", t.Val)
	// array = append(array, &t.Val)
	// fmt.Println("array", array)
	// walk the left side first
	walk(t.Left)

	// Inorder (Left, Root, Right)
	// fmt.Println(t.Val)
	// array = append(array, &t.Val)
	// fmt.Println("array", array)
	// walk the right side last
	walk(t.Right)

	// Postorder (Left, Right, Root)
	// fmt.Println("val", t.Val)
	// array = append(array, &t.Val)
	// fmt.Println("array", array)

}

func heightOfTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var left = heightOfTree(root.Left)
	var right = heightOfTree(root.Right)

	if left > right {
		return left + 1
	}

	return right + 1
}
