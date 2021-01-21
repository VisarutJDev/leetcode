package easy

// NumIdenticalPairs question
// Given an array of integers nums.
// A pair (i,j) is called good if nums[i] == nums[j] and i < j.
// Return the number of good pairs.
func NumIdenticalPairs(nums []int) int {
	l := len(nums)
	if l < 1 || l > 100 {
		return 0
	}

	i := 0
	j := i + 1

	pairCounter := 0
	for {
		if nums[i] < 1 || nums[i] > 100 {
			return 0
		}

		if nums[i] == nums[j] {
			pairCounter++
		}

		if j == (l - 1) {
			if i == (l - 2) {
				break
			}
			i++
			j = i + 1
			continue
		}

		j++
	}

	return pairCounter
}
