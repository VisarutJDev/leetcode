package easy

import (
	"fmt"
	"testing"
)

func TestInterpret(t *testing.T) {
	fmt.Println(Interpret("G()(al)"))
	fmt.Println(Interpret("G()()()()(al)"))
	fmt.Println(Interpret("(al)G(al)()()G"))
}
