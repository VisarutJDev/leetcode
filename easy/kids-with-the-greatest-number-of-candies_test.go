package easy

import (
	"fmt"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	fmt.Println(KidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
	fmt.Println(KidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
	fmt.Println(KidsWithCandies([]int{12, 1, 12}, 10))
}
