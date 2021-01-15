package easy

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	fmt.Println(RomanToInt("III"))
	fmt.Println(RomanToInt("IV"))
	fmt.Println(RomanToInt("IX"))
	fmt.Println(RomanToInt("LVIII"))
	fmt.Println(RomanToInt("MCMXCIV"))
}
