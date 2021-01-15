package easy

// IsPalindrome question
// Given an integer x, return true if x is palindrome integer.
// An integer is a palindrome when it reads the same backward as forward. For example, 121 is palindrome while 123 is not.
// the logic is to reverse number and check equal
// Note use logic from reverse-integer.go
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// x should be in range of -2^31 to 2^31 - 1
	if x > 2147483647 { //|| x < -2147483648 {
		return false
	}

	if x >= 0 && x <= 9 {
		return true
	}

	var reverse int = 0
	temp := x
	// reverse digit
	for x > 0 {
		float := x % 10  // use last digit
		reverse *= 10    // move digit
		reverse += float // assign new value to that digit
		x /= 10          // cut last digit to calculate next
	}

	if temp == reverse {
		return true
	}

	return false
}
