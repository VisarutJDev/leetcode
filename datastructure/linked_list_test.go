package datastructure

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := LinkedList{}
	node1 := &LinkedListNode{Data: 48}
	node2 := &LinkedListNode{Data: 18}
	node3 := &LinkedListNode{Data: 16}
	node4 := &LinkedListNode{Data: 11}
	node5 := &LinkedListNode{Data: 7}
	node6 := &LinkedListNode{Data: 2}
	list.Prepend(node1)
	list.Prepend(node2)
	list.Prepend(node3)
	list.Prepend(node4)
	list.Prepend(node5)
	list.Prepend(node6)
	list.PrintListData()
	list.DeleteWithValue(16)
	list.PrintListData()
}
