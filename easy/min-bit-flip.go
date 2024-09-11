package easy

import (
	"fmt"
)

func minBitFlips(start int, goal int) int {
	// change start and goal to binary
	// find difference in bit such as len and each bit to determine leading zero
	binStart := fmt.Sprintf("%b", start)
	binGoal := fmt.Sprintf("%b", goal)

	maxLen := len(binStart)

	if len(binGoal) > maxLen {
		maxLen = len(binGoal)
		binStart = fmt.Sprintf("%0*s", len(binGoal), binStart)
	} else {
		binGoal = fmt.Sprintf("%0*s", maxLen, binGoal)
	}

	rS := []rune(binStart)
	rG := []rune(binGoal)
	count := 0
	for i := maxLen - 1; i >= 0; i-- {
		if rS[i] != rG[i] {
			rS[i] = rG[i]
			count++
		}
	}
	return count
}
