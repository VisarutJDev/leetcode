package easy

import (
	"math/big"
	"strconv"
)

// PlusOne question
// Given a non-empty array of decimal digits representing a non-negative integer, increment one to the integer.
// The digits are stored such that the most significant digit is at the head of the list, and each element in the array contains a single digit.
// You may assume the integer does not contain any leading zero, except the number 0 itself.
func PlusOne(digits []int) []int {
	// brute force logic
	// i := len(digits) - 1
	// isPlus := false
	// for {
	// 	if i < 0 || isPlus {
	// 		break
	// 	}
	// 	if digits[i] == 9 {
	// 		digits[i] = 0
	// 		j := i - 1
	// 		if j >= 0 {
	// 			for {
	// 				if j == 0 {
	// 					if digits[j] == 9 {
	// 						digits[j] = 0
	// 						digits = append([]int{1}, digits...)
	// 					} else {
	// 						digits[j]++
	// 					}
	// 					isPlus = true
	// 					break
	// 				} else {

	// 				}
	// 			}
	// 		} else {
	// 			digits = append([]int{1}, digits...)
	// 			isPlus = true
	// 			break
	// 		}

	// 	} else {
	// 		digits[i]++
	// 		isPlus = true
	// 		break
	// 	}
	// 	i--
	// }
	// return digits

	// convert logic
	str := ""
	var result []int
	isLeadingZero := false
	for i := range digits {
		if i == 0 && digits[i] == 0 {
			isLeadingZero = true
		}
		str += strconv.Itoa(digits[i])
	}

	value := new(big.Int)
	value, _ = value.SetString(str, 10)
	value.Add(value, big.NewInt(1))

	newStr := value.String()
	for i := range newStr {
		value, _ := strconv.Atoi(string(newStr[i]))
		result = append(result, value)
	}

	if isLeadingZero && len(str) != len(result) {
		extend := make([]int, len(str)-len(result))
		result = append(extend, result...)
	}
	return result
}
