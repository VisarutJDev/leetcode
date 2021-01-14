package easy

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	fmt.Println(Reverse(123))
	fmt.Println(Reverse(-123))
	fmt.Println(Reverse(120))
	fmt.Println(Reverse(0))
}
