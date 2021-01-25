package easy

import (
	"strconv"
)

// CountAndSay question
// The count-and-say sequence is a sequence of digit strings defined by the recursive formula:
// countAndSay(1) = "1"
// countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1), which is then converted into a different digit string.
// To determine how you "say" a digit string, split it into the minimal number of groups so that each group is a contiguous section all of the same character.
// Then for each group, say the number of characters, then say the character.
// To convert the saying into a digit string, replace the counts with a number and concatenate every saying.
// For example, the saying and conversion for digit string "3322251":
func CountAndSay(n int) string {
	if n < 1 || n > 30 {
		return ""
	}

	result := "1"
	if n == 1 {
		return result
	}
	for i := 1; i < n; i++ {
		strCount := ""
		if i > n {
			break
		}
		count := 1
		for j := 1; j < len(result); j++ {
			// if current eq previous means that it two {that number}. Ex if n = 3 it will be 1,11 (two 1) so next is 21
			if result[j] == result[j-1] {
				count++
			} else {
				strCount += strconv.Itoa(count)
				strCount += string(result[j-1])
				count = 1
			}
		}
		strCount += strconv.Itoa(count)
		strCount += string(result[len(result)-1])
		result = strCount
	}

	return result
}
