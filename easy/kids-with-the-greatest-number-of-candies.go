package easy

// KidsWithCandies question
// Given the array candies and the integer extraCandies, where candies[i] represents the number of candies that the ith kid has.
// For each kid check if there is a way to distribute extraCandies among the kids such that he or she can have the greatest number of candies among them.
// Notice that multiple kids can have the greatest number of candies.
func KidsWithCandies(candies []int, extraCandies int) []bool {
	numOfKids := len(candies)
	result := make([]bool, numOfKids)
	maxVal := max(candies)
	for i := range candies {
		n := candies[i] + extraCandies
		result[i] = (n >= maxVal)
	}
	return result
}

func max(array []int) int {
	var max int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}
