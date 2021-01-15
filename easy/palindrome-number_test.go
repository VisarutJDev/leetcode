package easy

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	fmt.Println(IsPalindrome(121))
	fmt.Println(IsPalindrome(-121))
	fmt.Println(IsPalindrome(10))
	fmt.Println(IsPalindrome(-101))
}
