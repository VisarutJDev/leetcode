package easy

// RemoveDuplicates question
// Given a sorted array nums, remove the duplicates in-place such that each element appears only once and returns the new length.
// Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
// Clarification:
// Confused why the returned value is an integer but your answer is an array?
// Note that the input array is passed in by reference, which means a modification to the input array will be known to the caller as well.
func RemoveDuplicates(nums []int) int {

	if len(nums) <= 1 || len(nums) > 30000 {
		return len(nums)
	}

	i := 0
	j := 1
	for {

		if len(nums) <= 1 || j > (len(nums)-1) {
			return len(nums)
		}

		if nums[i] < -10000 || nums[i] > 10000 {
			return 0
		}

		if nums[i] == nums[j] {
			nums = remove(nums, j)
		} else {
			if j == (len(nums) - 1) {
				if i == (len(nums) - 1) {
					break
				}

				i++
				j = i + 1
				if j > (len(nums) - 1) {
					break
				}
				continue
			}
			j++
		}
	}

	return len(nums)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
