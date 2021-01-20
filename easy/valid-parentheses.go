package easy

// IsValid question
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
//   1. Open brackets must be closed by the same type of brackets.
//   2. Open brackets must be closed in the correct orde
func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var lastestOpenBrak []string

	for i := range s {
		if i == 0 && !(string(s[i]) == "(" || string(s[i]) == "{" || string(s[i]) == "[") {
			return false
		}

		if string(s[i]) == "(" || string(s[i]) == "{" || string(s[i]) == "[" {
			lastestOpenBrak = append(lastestOpenBrak, string(s[i]))
		}

		switch string(s[i]) {
		case ")":
			if len(lastestOpenBrak) != 0 {
				if lastestOpenBrak[len(lastestOpenBrak)-1] != "(" {
					return false
				}
				lastestOpenBrak = lastestOpenBrak[:len(lastestOpenBrak)-1]
			} else {
				return false
			}
		case "}":
			if len(lastestOpenBrak) != 0 {
				if lastestOpenBrak[len(lastestOpenBrak)-1] != "{" {
					return false
				}
				lastestOpenBrak = lastestOpenBrak[:len(lastestOpenBrak)-1]
			} else {
				return false
			}
		case "]":
			if len(lastestOpenBrak) != 0 {
				if lastestOpenBrak[len(lastestOpenBrak)-1] != "[" {
					return false
				}
				lastestOpenBrak = lastestOpenBrak[:len(lastestOpenBrak)-1]
			} else {
				return false
			}
		}

	}

	if len(lastestOpenBrak) != 0 {
		return false
	}

	return true
}
