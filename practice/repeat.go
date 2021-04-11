package practice

import (
	"fmt"
	"strings"
)

func CountRepeatCharactor(s string) string {
	var b strings.Builder
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			b.WriteString(fmt.Sprintf("%s%d", string(s[i-1]), count))
			count = 1
			continue
		}
		count++
		if i == len(s)-1 {
			b.WriteString(fmt.Sprintf("%s%d", string(s[i-1]), count))
		}
	}

	if len(b.String()) == len(s) {
		return s
	}
	return b.String()
}
