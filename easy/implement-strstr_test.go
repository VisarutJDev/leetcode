package easy

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	fmt.Println(StrStr("hello", "ll"))
	fmt.Println(StrStr("aaaaa", "bba"))
	fmt.Println(StrStr("", ""))
	fmt.Println(StrStr("", "a"))
}
