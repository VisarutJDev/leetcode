package easy

// RemoveElement question
// Given an array nums and a value val, remove all instances of that value in-place and return the new length.
// Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
// The order of elements can be changed. It doesn't matter what you leave beyond the new length.
// Clarification:
// Confused why the returned value is an integer but your answer is an array?
// Note that the input array is passed in by reference, which means a modification to the input array will be known to the caller as well.
func RemoveElement(nums []int, val int) int {

	if len(nums) <= 0 || len(nums) > 100 {
		return len(nums)
	}

	if val < 0 || val > 100 {
		return len(nums)
	}

	i := 0

	for {
		if i > len(nums)-1 {
			break
		}

		if nums[i] < 0 {
			return len(nums)
		}

		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			if i == len(nums)-1 {
				break
			}
			i++
		}

	}

	return len(nums)
}
