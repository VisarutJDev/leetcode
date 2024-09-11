package easy

import (
	"fmt"
	"testing"
)

func TestMinBitFlips(t *testing.T) {
	// change start and goal to binary
	// find difference in bit such as len and each bit to determine leading zero
	fmt.Println(minBitFlips(10, 7))
	fmt.Println(minBitFlips(3, 4))
}
