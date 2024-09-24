package hackerrank

import "fmt"

func hourglassSum(arr [][]int32) int32 {
	// Write your code here
	// hour glass pattern
	// 00 01 02 >> ij    ij+1  ij+2
	//    11    >>      i+1j+1
	// 20 21 22 >> i+2j i+2j+1 i+2j+2
	// check that j+2 not more than len(arr[i])
	// check that i+2 not more than len(arr[i]) , because same dimention 6*6
	var maxSum int32 = -9999
	if len(arr) < 3 {
		// not possible in hourglass form
		return 0
	}
	i, j := 0, 0
	for i+2 < len(arr[i]) {
		fmt.Println(arr[i][j], arr[i][j+1], arr[i][j+2])
		fmt.Println(arr[i+1][j+1])
		fmt.Println(arr[i+2][j], arr[i+2][j+1], arr[i+2][j+2])
		sum := arr[i][j] + arr[i][j+1] + arr[i][j+2]
		sum += arr[i+1][j+1]
		sum += arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
		fmt.Println(sum, maxSum)
		if sum > maxSum {
			maxSum = sum
		}
		j++
		if j+2 >= len(arr[i]) {
			// reset j and inc i
			j = 0
			i++
		}

	}

	return maxSum
}
