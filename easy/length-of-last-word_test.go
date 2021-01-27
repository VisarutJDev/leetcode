package easy

import (
	"fmt"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	fmt.Println(LengthOfLastWord("Hello World"))
	fmt.Println(LengthOfLastWord(" "))
	fmt.Println(LengthOfLastWord("a "))
}
