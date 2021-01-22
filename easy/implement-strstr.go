package easy

import (
	"strings"
)

// StrStr question
// Implement strStr().
// Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
// Clarification:
// What should we return when needle is an empty string? This is a great question to ask during an interview.
// For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
func StrStr(haystack string, needle string) int {
	if len(haystack) > 50000 || len(haystack) < 0 || len(needle) > 50000 || len(needle) < 0 {
		return 0
	}

	// not contain
	return strings.Index(haystack, needle)
}
