package easy

// NumJewelsInStones question
// You're given strings jewels representing the types of stones that are jewels,
// and stones representing the stones you have. Each character in stones is a type of stone you have.
// You want to know how many of the stones you have are also jewels.
// Letters are case sensitive, so "a" is considered a different type of stone from "A".
func NumJewelsInStones(jewels string, stones string) int {

	if len(jewels) < 1 || len(jewels) > 50 || len(stones) < 1 || len(stones) > 50 {
		return 0
	}

	i := 0
	j := 0
	counter := 0
	for {

		if string(jewels[i]) == string(stones[j]) {
			counter++
		}

		if j == len(stones)-1 {
			i++
			j = 0
		} else {
			j++
		}

		if i > len(jewels)-1 {
			break
		}

	}
	return counter
}
