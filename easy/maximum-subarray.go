package easy

// MaxSubArray question
// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
func MaxSubArray(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	isFoundPositive := false
	maxNegative := -100000

	maxSofar := 0
	maxEnding := 0

	for i := range nums {
		if nums[i] > 0 {
			isFoundPositive = true
		}
		if nums[i] > maxNegative {
			maxNegative = nums[i]
		}

		maxEnding = maxEnding + nums[i]
		if maxSofar < maxEnding {
			maxSofar = maxEnding
		}

		if maxEnding < 0 {
			maxEnding = 0
		}
	}

	if !isFoundPositive {
		return maxNegative
	}

	return maxSofar
}
