package easy

// SearchInsert question
// Given a sorted array of distinct integers and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.
func SearchInsert(nums []int, target int) int {
	if len(nums) > 10000 || len(nums) < 1 {
		return 0
	}

	if target > 10000 || target < -10000 {
		return 0
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > 10000 || nums[i] < -10000 {
			return 0
		}

		if nums[i] == target {
			return i
		}

		if nums[i] < target {
			return i + 1
		}

	}

	return 0
}
