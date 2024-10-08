package hackerrank

import (
	"fmt"
	"testing"
)

func TestHourglassSum(t *testing.T) {
	fmt.Println(hourglassSum(
		[][]int32{
			{0, -4, -6, 0, -7, -6},
			{-1, -2, -6, -8, -3, -1},
			{-8, -4, -2, -8, -8, -6},
			{-3, -1, -2, -5, -7, -4},
			{-3, -5, -3, -6, -6, -6},
			{-3, -6, 0, -8, -6, -7},
		},
	))
}
