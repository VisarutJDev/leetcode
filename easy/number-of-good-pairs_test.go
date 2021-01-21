package easy

import (
	"fmt"
	"testing"
)

func TestNumIdenticalPairs(t *testing.T) {

	fmt.Println(NumIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
	fmt.Println(NumIdenticalPairs([]int{1, 1, 1, 1}))

	fmt.Println(NumIdenticalPairs([]int{1, 2, 3}))

}
