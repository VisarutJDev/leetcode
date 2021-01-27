package easy

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	fmt.Println(MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(MaxSubArray([]int{1}))
	fmt.Println(MaxSubArray([]int{0}))
	fmt.Println(MaxSubArray([]int{-1}))
	fmt.Println(MaxSubArray([]int{-100000}))
	fmt.Println(MaxSubArray([]int{-2, -1}))
	fmt.Println(MaxSubArray([]int{1, 2, -1}))
}
