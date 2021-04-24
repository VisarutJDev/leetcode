package practice

// Create a function to create binary search tree from increasing order number
func CreateBST(number []int) *Node {
	// find middle because increasing order number mid will be middle value
	mid := (len(number) - 1) / 2
	headNode := &Node{val: number[mid]}
	number = remove(number, mid)
	currentNode := headNode
	for i := range number {
		// fmt.Println(number[i])
		currentNode.Insert(number[i])
		currentNode = headNode
	}

	return headNode
}

// Insert to insert
func (n *Node) Insert(number int) {
	// fmt.Println(n.val)
	// fmt.Println(number)
	if number > n.val {
		// move right
		if n.right == nil {
			n.right = &Node{val: number}
		} else {
			n.right.Insert(number)
		}
	} else if number < n.val {
		//move left
		if n.left == nil {
			n.left = &Node{val: number}
		} else {
			n.left.Insert(number)
		}
	}
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

// Create a function to create binary search tree from increasing order number
// แบ่งครึ่ง เพื่อเอา node root ไปเรื่อยๆ ซึ่ง จะถูก
func CreateBSTV2(number []int, start int, end int) *Node {
	if start > end {
		return nil
	}
	// find middle because increasing order number mid will be middle value
	mid := (start + end) / 2
	root := &Node{val: number[mid]}
	root.left = CreateBSTV2(number, start, mid-1)
	root.right = CreateBSTV2(number, mid+1, end)

	return root
}

// LinkedDepth is element to hold node int
type LinkedDepth struct {
	Data int
	Next *LinkedDepth
}

//CreateLevelLinkedList โดยการย้ายไป node ขวาทั้งหมด หรือซ้าย ก็จะกลายเป็น linked list ได้
func CreateLevelLinkedList(root *Node) {
	if root == nil {
		return
	}

	// leaf node or last element
	if root.left == nil && root.right == nil {
		return
	}

	if root.left != nil {
		CreateLevelLinkedList(root.left)

		temp := root.right
		root.right = root.left
		root.left = nil

		current := root.right
		for current.right != nil {
			current = current.right
		}
		current.right = temp
	}

	CreateLevelLinkedList(root.right)
}
