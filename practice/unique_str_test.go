package practice

import (
	"fmt"
	"testing"
)

func TestCheckAllUniqueString(t *testing.T) {
	fmt.Println(CheckAllUniqueString("John"))
	fmt.Println(CheckAllUniqueString("JohnSmith"))
	fmt.Println(CheckAllUniqueString(""))
	fmt.Println(CheckAllUniqueString("J"))
	fmt.Println(CheckAllUniqueString("JJ"))
}
