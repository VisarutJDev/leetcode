package hackerrank

import "strings"

func appendAndDelete(s string, t string, k int32) string {
	// Write your code here
	// if operation is under k still Yes? because we can use deletion on empty string to get empty string
	var countOperation int32 = 0
	// countOperation < k Yes
	// this mean we have to delete some
	// >> some of first string contains in second string
	// this mean we have to delete it all
	// >> non in first string contains in second string
	if s == t && len(s) == len(t) {
		return "Yes"
	}

	for k > 0 {
		if countOperation > k {
			return "No"
		}
		if strings.Contains(t, s) {
			restOfTheWord := t[len(s):]
			if int32(len(restOfTheWord))+countOperation > k {
				return "No"
			}
			if s+restOfTheWord == t && int32(len(restOfTheWord))+countOperation == k {
				return "Yes"
			}
		}
		if len(s) > 0 {
			s = s[:len(s)-1]
		}
		countOperation++
	}
	return "No"
}
