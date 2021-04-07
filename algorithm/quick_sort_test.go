package algorithm

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	test := []int{2, 6, 5, 3, 8, 7, 1, 0, 4}
	fmt.Println(test)
	fmt.Println(QuickSort(test))
}
