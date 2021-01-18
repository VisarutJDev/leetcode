package easy

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	fmt.Println(LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(LongestCommonPrefix([]string{"dog", "racecar", "car"}))
}
