package easy

// Reverse question
// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
// Logic is to calculate from last digit of int by modulate, and multiply 10 to move digit from unit => tenth => hundred => so on,
// then use int that modulate concatenate result
func Reverse(x int) int {
	result := 0
	input := x
	if input < 0 {
		input *= -1
	}
	for input > 0 {
		float := input % 10 // use last digit
		result *= 10        // move digit
		result += float     // assign new value to that digit
		input /= 10         // cut last digit to calculate next
	}
	if x < 0 {
		result *= -1
	}
	// result more than 2^31 - 1 or less than -2^31
	if result > 2147483647 || result < -2147483648 {
		return 0
	}
	return result
}
