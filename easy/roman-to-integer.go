package easy

// RomanToInt question
// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
//
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
//
// For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII,
// which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
// Roman numerals are usually written largest to smallest from left to right. However,
// the numeral for four is not IIII. Instead, the number four is written as IV.
// Because the one is before the five we subtract it making four.
// The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
//
// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
func RomanToInt(s string) int {
	Roman := make(map[string]int)
	Roman["I"] = 1
	Roman["V"] = 5
	Roman["X"] = 10
	Roman["L"] = 50
	Roman["C"] = 100
	Roman["D"] = 500
	Roman["M"] = 1000

	total := 0
	skipNext := false
	for i := range s {
		if skipNext {
			skipNext = false
			continue
		}

		if i != len(s)-1 {
			// check if following of I have V or X
			if string(s[i]) == "I" && (string(s[i+1]) == "V" || string(s[i+1]) == "X") {
				skipNext = true
				total += Roman[string(s[i+1])] - Roman[string(s[i])]
				continue
			}
			// check if following of X have L or C
			if string(s[i]) == "X" && (string(s[i+1]) == "L" || string(s[i+1]) == "C") {
				skipNext = true
				total += Roman[string(s[i+1])] - Roman[string(s[i])]
				continue
			}
			// check if following of C have D or M
			if string(s[i]) == "C" && (string(s[i+1]) == "D" || string(s[i+1]) == "M") {
				skipNext = true
				total += Roman[string(s[i+1])] - Roman[string(s[i])]
				continue
			}
		}

		total += Roman[string(s[i])]
	}

	return total
}
