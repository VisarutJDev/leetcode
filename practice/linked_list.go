package practice

import (
	"fmt"
)

// NodeString is element for linked link node string
type NodeString struct {
	Data string
	Next *NodeString
}

// NodeInt is element to hold node int
type NodeInt struct {
	Data int
	Next *NodeInt
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

// GetValue is a function to get value from node int
func GetValue(l *NodeInt) int {
	unit := 10
	c := l
	sum := c.Data
	for c.Next != nil {
		sum += c.Next.Data * unit
		unit *= 10
		c = c.Next
	}
	return sum
}

// Add is a function to add 2 node int
func Add(l1, l2 *NodeInt) int {
	return GetValue(l1) + GetValue(l2)
}

// AppendNode to append node int
func (n *NodeInt) AppendNode(s int) {
	newNode := &NodeInt{Data: s}
	currentNode := n
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = newNode
}

// GetValueRe is a function to get value from node int
func GetValueRe(l *NodeInt) int {
	var arr []int
	unit := 1
	c := l
	arr = append(arr, c.Data)
	sum := 0
	for c.Next != nil {
		arr = append(arr, c.Next.Data)
		unit *= 10
		c = c.Next
	}
	for i := range arr {
		sum += arr[i] * unit
		unit /= 10
	}
	return sum
}

// AddRe is a function to addRe 2 node int
func AddRe(l1, l2 *NodeInt) int {
	return GetValueRe(l1) + GetValueRe(l2)
}

func (n *NodeString) CheckNodeLoop() string {
	check := make(map[string]bool)
	check[n.Data] = true
	c := n
	for c.Next != nil {
		if check[c.Next.Data] {
			return c.Next.Data
		}
		check[c.Next.Data] = true
		c = c.Next
	}
	return ""
}
