package practice

import "fmt"

// NodeString is element for linked link node string
type NodeString struct {
	Data string
	Next *NodeString
}

// AppendNode to append node string
func (n *NodeString) AppendNode(s string) {
	newNode := &NodeString{Data: s}
	currentNode := n
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = newNode
}

// PrintNode to append node string
func (n *NodeString) PrintNode() {
	currentNode := n
	for currentNode.Next != nil {
		fmt.Println(currentNode)
		currentNode = currentNode.Next
	}
	fmt.Println(currentNode)
}

// RemoveDup to remove duplication charactor
func (n *NodeString) RemoveDup() {
	current := n
	check := make(map[string]bool)
	check[current.Data] = true

	for current.Next != nil {
		if check[current.Next.Data] {
			current.Next = current.Next.Next
			continue
		}
		check[current.Next.Data] = true
		current = current.Next
	}
}

//KthToElement to find kth to last element of linked list
func (n *NodeString) KthToElement(k int) []string {
	var arr []string
	c := n
	arr = append(arr, c.Data)
	for c.Next != nil {
		arr = append(arr, c.Next.Data)
		c = c.Next
	}

	if k-1 > len(arr) {
		return []string{}
	}

	return arr[k-1:]
}
