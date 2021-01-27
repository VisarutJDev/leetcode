package easy

import (
	"strings"
)

// LengthOfLastWord question
// Given a string s consists of some words separated by spaces, return the length of the last word in the string. If the last word does not exist, return 0.
// A word is a maximal substring consisting of non-space characters only.
func LengthOfLastWord(s string) int {
	strArr := strings.Split(s, " ")
	i := 1
	l := 0
	for {
		if i > len(strArr) {
			break
		}
		if len(strArr[len(strArr)-i]) == 0 {
			i++
		} else {
			l = len(strArr[len(strArr)-i])
			break
		}
	}

	return l
}
