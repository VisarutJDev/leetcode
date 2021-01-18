package easy

import (
	"fmt"
	"testing"
)

func TestShuffle(t *testing.T) {
	fmt.Println(Shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
	fmt.Println(Shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
	fmt.Println(Shuffle([]int{1, 1, 2, 2}, 2))
}
