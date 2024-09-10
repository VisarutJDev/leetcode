package hackerrank

import "fmt"

func ClimbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	rankingLadderPosition := []int32{}
	m := make(map[int32]bool, 0)
	arr := []int32{}
	for i := range ranked {
		if !m[ranked[i]] {
			arr = append(arr, ranked[i])
		}
		m[ranked[i]] = true
	}
	for j := 0; j < len(player); j++ {
		newarr := append(arr, player[j])
		QuickSort(newarr, 0, int32(len(newarr)-1))
		for k := 0; k < len(newarr); k++ {
			if newarr[k] == player[j] {
				rankingLadderPosition = append(rankingLadderPosition, int32(k+1))
				break
			}
		}
	}
	fmt.Println(rankingLadderPosition)
	return rankingLadderPosition
}

func ClimbingLeaderboard1(ranked []int32, player []int32) []int32 {
	// Write your code here
	rankingLadderPosition := []int32{}
	m := make(map[int32]bool, 0)
	arr := []int32{}
	for i := range ranked {
		if !m[ranked[i]] {
			arr = append(arr, ranked[i])
		}
		m[ranked[i]] = true
	}
	for j := 0; j < len(player); j++ {
		for i := len(arr) - 1; i >= 0; i-- {
			if i == 0 && player[j] > arr[i] {
				rankingLadderPosition = append(rankingLadderPosition, int32(i+1))
			}
			if player[j] < arr[i] {
				rankingLadderPosition = append(rankingLadderPosition, int32(i+2))
				break
			}
		}
	}
	return rankingLadderPosition
}

func Partition(arr []int32, low, high int32) int32 {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] >= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // Swap
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // Swap
	return i + 1
}

func QuickSort(arr []int32, low, high int32) {
	if low < high {
		pi := Partition(arr, low, high)
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

// BSTNode represents a node in a binary search tree
type BSTNode struct {
	Value int32
	Left  *BSTNode
	Right *BSTNode
}

// Insert inserts a value into the binary search tree
func (node *BSTNode) Insert(value int32) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &BSTNode{Value: value}
		} else {
			node.Left.Insert(value)
		}
	} else {
		if node.Right == nil {
			node.Right = &BSTNode{Value: value}
		} else {
			node.Right.Insert(value)
		}
	}
}

// Search searches for a value in the binary search tree
func (node *BSTNode) Search(value int32) bool {
	if node == nil {
		return false
	}
	if value < node.Value {
		return node.Left.Search(value)
	} else if value > node.Value {
		return node.Right.Search(value)
	}
	return true
}

func BFS(root *BSTNode) {
	if root == nil {
		return
	}

	queue := []*BSTNode{root}

	for len(queue) > 0 {
		// Dequeue the front node
		current := queue[0]
		queue = queue[1:]

		// Process the current node
		fmt.Printf("%d ", current.Value)

		// Enqueue left child
		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		// Enqueue right child
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}
