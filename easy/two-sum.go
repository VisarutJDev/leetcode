package easy

// TwoSum question
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
func TwoSum(nums []int, target int) []int {
	var result []int
	if len(nums) <= 1 {
		return result
	}
	for i := range nums {
		for j := 1; j < len(nums); j++ {
			if i == j {
				continue
			}
			if target == nums[i]+nums[j] {
				result = append(result, i)
				result = append(result, j)
				break
			}
		}
		if len(result) != 0 {
			break
		}
	}
	return result
}
