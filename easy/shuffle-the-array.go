package easy

// Shuffle question
// Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].
// Return the array in the form [x1,y1,x2,y2,...,xn,yn]
func Shuffle(nums []int, n int) []int {
	if n < 1 || n > 500 {
		return []int{}
	}

	if len(nums) != 2*n {
		return []int{}
	}

	var result []int
	for i := 0; i < n; i++ {
		if nums[i] < 1 || nums[i] > 1000 || nums[i+n] < 1 || nums[i+n] > 1000 {
			return []int{}
		}

		result = append(result, nums[i])
		result = append(result, nums[i+n])
	}

	return result
}
