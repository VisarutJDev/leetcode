package easy

import (
	"fmt"
	"testing"
)

func TestPlusOne(t *testing.T) {
	fmt.Println(PlusOne([]int{1, 2, 3}))
	fmt.Println(PlusOne([]int{4, 3, 2, 1}))
	fmt.Println(PlusOne([]int{0}))
	fmt.Println(PlusOne([]int{9}))
	fmt.Println(PlusOne([]int{9, 9}))
	fmt.Println(PlusOne([]int{0, 0}))
	fmt.Println(PlusOne([]int{1, 0}))
	fmt.Println(PlusOne([]int{0, 0, 0}))
	fmt.Println(PlusOne([]int{0, 9}))
	fmt.Println(PlusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}))
}
