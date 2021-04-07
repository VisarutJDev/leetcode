package algorithm

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	fmt.Println("\n--- Sorted ---\n\n", mergeSort(slice))
}
