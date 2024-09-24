package hackerrank

func timeInWords(h int32, m int32) string {
	// Write your code here
	// 0 o'clock
	// 1<= minute <= 30 past
	// minute > 30 to
	result := ""
	specialCase := map[int32]string{
		0:  "o' clock",
		1:  "one minute",
		15: "quarter",
		30: "haft",
		45: "quarter",
	}
	numbers := map[int32]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		// 20: "twenty",
	}

	concatTxt := " "

	if m >= 1 && m <= 30 {
		concatTxt = " past "
	} else if m > 30 {
		concatTxt = " to "
	}
	switch m {
	case 0:
		result += numbers[h] + concatTxt + specialCase[m]
		return result
	case 1, 15, 30:
		result += specialCase[m] + concatTxt + numbers[h]
		return result
	case 45:
		result += specialCase[m] + concatTxt + numbers[h+1]
		return result
	}
	concatTxt = " minutes" + concatTxt
	switch m / 10 {
	case 0:
		result += numbers[m] + concatTxt + numbers[h]
		return result
	case 1, 2:
		digit := m % 10
		if digit >= 0 && digit <= 3 {
			result += numbers[m] + concatTxt + numbers[h]
			return result
		}
		if m/10 == 1 {
			result += numbers[digit] + "teen " + concatTxt + numbers[h]
		} else {
			result += "twenty " + numbers[digit] + concatTxt + numbers[h]
		}
		return result
	case 3, 4, 5:
		minuteLeft := 60 - m
		digit := minuteLeft % 10
		if digit >= 0 && digit <= 3 {
			result += numbers[minuteLeft] + concatTxt + numbers[h+1]
			return result
		}
		if minuteLeft/10 == 1 {
			result += numbers[digit] + "teen " + concatTxt + numbers[h+1]
		} else {
			result += "twenty " + numbers[digit] + concatTxt + numbers[h+1]
		}
		return result
	}

	return result
}
