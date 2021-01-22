package easy

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	fmt.Println(SearchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(SearchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(SearchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(SearchInsert([]int{1, 3, 5, 6}, 0))
	fmt.Println(SearchInsert([]int{1}, 0))
}
