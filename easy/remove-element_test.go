package easy

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	fmt.Println(RemoveElement([]int{1}, 1))
	fmt.Println(RemoveElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
