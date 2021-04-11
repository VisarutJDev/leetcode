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
