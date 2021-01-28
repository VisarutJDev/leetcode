package easy

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	fmt.Println(Decode([]int{1, 2, 3}, 1))
	fmt.Println(Decode([]int{6, 2, 7, 3}, 4))
}
