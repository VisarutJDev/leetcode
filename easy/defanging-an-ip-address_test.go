package easy

import (
	"fmt"
	"testing"
)

func TestDefangIPaddr(t *testing.T) {
	fmt.Println(DefangIPaddr("1.1.1.1"))
	fmt.Println(DefangIPaddr("255.100.50.0"))
}
