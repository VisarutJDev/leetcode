package easy

import (
	"fmt"
	"math/big"
)

// AddBinary question
// Given two binary strings a and b, return their sum as a binary string.
func AddBinary(a string, b string) string {
	value := new(big.Int)
	value, _ = value.SetString(a, 2)

	value2 := new(big.Int)
	value2, _ = value2.SetString(b, 2)

	value.Add(value, value2)

	return fmt.Sprintf("%b", value)
}
