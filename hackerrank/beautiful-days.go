package hackerrank

func BeautifulDays(i int32, j int32, k int32) int32 {
	// Write your code here
	// 1. reverse number of i to j
	// 2. i*reverseNumber(i) % 6 if it more than 0 means it's not BeautifulDays
	// 3. if it 0 mean BeautifulDays
	var count int32 = 0
	for i <= j {
		if (i-reverseNumber(i))%k == 0 {
			count++
		}
		i++
	}
	return count
}

func reverseNumber(numberToReverse int32) int32 {
	var result int32 = 0
	input := numberToReverse

	// Constraints show that no negative value below is no need
	// if input < 0 {
	// 	input *= -1
	// }

	for input > 0 {
		lastNumberinNumber := input % 10
		// at first result 0 so this mean notthing at first
		result *= 10
		// just like append to array
		result += lastNumberinNumber
		// remove lastNumberinNumber and continue to next number
		input /= 10
	}

	return result
}
