package easy

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	fmt.Println(TwoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(TwoSum([]int{3, 2, 4}, 6))
	fmt.Println(TwoSum([]int{3, 3}, 6))
}
