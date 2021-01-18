package easy

import (
	"sort"
)

// LongestCommonPrefix question
// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".
func LongestCommonPrefix(strs []string) string {
	n := len(strs)
	if n > 200 || n <= 0 {
		return ""
	}

	for i := 0; i < n; i++ {
		if len(strs[i]) < 0 || len(strs[i]) > 200 {
			return ""
		}
	}

	sort.Strings(strs)

	en := 0
	if len(strs[0]) < len(strs[n-1]) {
		en = len(strs[0])
	} else {
		en = len(strs[n-1])
	}

	first, last := strs[0], strs[n-1]

	i := 0
	for i < en && first[i] == last[i] {
		i++
	}

	return first[:i]
}
