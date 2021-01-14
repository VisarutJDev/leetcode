package easy

import (
	"fmt"
	"testing"
)

func TestMaximumWealth(t *testing.T) {
	input := [][]int{
		[]int{1, 2, 3},
		[]int{3, 2, 1},
	}
	fmt.Println(MaximumWealth(input))
	input = [][]int{
		[]int{1, 5},
		[]int{7, 3},
		[]int{3, 5},
	}
	fmt.Println(MaximumWealth(input))
	input = [][]int{
		[]int{2, 8, 7},
		[]int{7, 1, 3},
		[]int{1, 9, 5},
	}
	fmt.Println(MaximumWealth(input))
}
