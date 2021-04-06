package datastructure

import "fmt"

// LinkedListNode represent node of linked list
type LinkedListNode struct {
	Data int
	Next *LinkedListNode
}

// LinkedList represent list
type LinkedList struct {
	Head   *LinkedListNode
	Length int
}

func (l *LinkedList) Prepend(n *LinkedListNode) {
	secondHead := l.Head
	l.Head = n
	l.Head.Next = secondHead
	l.Length++
}

func (l LinkedList) PrintListData() {
	toPrint := l.Head
	length := l.Length
	for length != 0 {
		fmt.Printf("%d ", toPrint.Data)
		toPrint = toPrint.Next
		length--
	}
	fmt.Println()
}

func (l *LinkedList) DeleteWithValue(value int) {
	if l.Length == 0 {
		return
	}

	if l.Head.Data == value {
		l.Head = l.Head.Next
		l.Length--
		return
	}

	// it's hard to modified next node from previous node so we will compare next node
	toDelete := l.Head
	for toDelete.Next.Data != value {
		if toDelete.Next.Next == nil {
			return
		}
		toDelete = toDelete.Next
	}
	toDelete.Next = toDelete.Next.Next
	l.Length--
}
