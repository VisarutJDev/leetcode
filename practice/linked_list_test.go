package practice

import (
	"fmt"
	"testing"
)

func TestRemoveDup(t *testing.T) {
	node := &NodeString{Data: "F"}
	// node.AppendNode("F")
	node.AppendNode("O")
	node.AppendNode("L")
	node.AppendNode("L")
	node.AppendNode("O")
	node.AppendNode("W")
	node.AppendNode(" ")
	node.AppendNode("U")
	node.AppendNode("P")
	node.PrintNode()
	node.RemoveDup()
	node.PrintNode()
	// fmt.Println(node.KthToElement(7))
	fmt.Println(node.KthToElement(3))
}

func TestAdd(t *testing.T) {
	node1 := &NodeInt{Data: 7}
	node1.AppendNode(1)
	node1.AppendNode(6)
	node2 := &NodeInt{Data: 5}
	node2.AppendNode(9)
	node2.AppendNode(2)
	fmt.Println(Add(node1, node2))
}
func TestAddRe(t *testing.T) {
	node1 := &NodeInt{Data: 6}
	node1.AppendNode(1)
	node1.AppendNode(7)
	node2 := &NodeInt{Data: 2}
	node2.AppendNode(9)
	node2.AppendNode(5)
	fmt.Println(AddRe(node1, node2))
}

func TestCheckNodeLoop(t *testing.T) {
	node := &NodeString{Data: "A"}
	node.AppendNode("B")
	node.AppendNode("C")
	node.AppendNode("D")
	node.AppendNode("E")
	node.AppendNode("C")
	fmt.Println(node.CheckNodeLoop())
}
