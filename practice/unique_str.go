package practice

// CheckAllUniqueString is a function to check string input is unique in each charactor
// True is unique
// False is not
func CheckAllUniqueString(s string) bool {
	hashmap := make(map[rune]bool)
	for _, c := range s {
		if hashmap[c] {
			return false
		}
		hashmap[c] = true
	}
	return true
}
