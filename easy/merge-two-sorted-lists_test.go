package easy

import (
	"fmt"
	"testing"
)

func newMergeLinkedList(n int) *ListNode { return &ListNode{Val: n} }

func TestMergeTwoLists(t *testing.T) {
	head := newMergeLinkedList(1)
	head.Next = newMergeLinkedList(2)
	head.Next.Next = newMergeLinkedList(4)

	head2 := newMergeLinkedList(1)
	head2.Next = newMergeLinkedList(3)
	head2.Next.Next = newMergeLinkedList(4)

	fmt.Println(MergeTwoLists(head, head2))

	head = nil
	head2 = nil
	fmt.Println(MergeTwoLists(head, head2))

	head = nil
	head2 = newMergeLinkedList(0)
	fmt.Println(MergeTwoLists(head, head2))
}
