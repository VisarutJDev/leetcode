package easy

import "fmt"

// ListNode is a list node
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoLists question
// Merge two sorted linked lists and return it as a sorted list. The list should be made by splicing together the nodes of the first two lists.
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var queue *ListNode
	// to create first node to list
	firstQ := &ListNode{Val: -1, Next: nil}
	queue = firstQ
	for {
		if l1 != nil && l2 != nil {
			fmt.Println("loop =>", l1)
			fmt.Println("loop =>", l2)
			// check whenever element is less than other element
			// which one is less put in to queue
			// and move to next element for list that is less
			if l1.Val <= l2.Val {
				firstQ.Next = l1
				l1 = l1.Next
			} else {
				firstQ.Next = l2
				l2 = l2.Next
			}

			// move to next node
			// next node already assign from above
			firstQ = firstQ.Next
		} else {
			break
		}

	}

	// set last element
	// case list is not nil means that it is last element
	if l1 != nil {
		firstQ.Next = l1
	}
	if l2 != nil {
		firstQ.Next = l2
	}

	// becuase first q is initial value
	return queue.Next
}
