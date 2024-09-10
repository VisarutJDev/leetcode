package hackerrank

func AngryProfessor(k int32, a []int32) string {
	// Write your code here
	// 1. sort an array input because we should count only late student and ignore on time student
	// 2. after we sort we will have positive numbers (late student) to count and then easily stop when reach threshold
	sorted := mSort(a)
	var count int32 = 0
	for i := 0; i < len(sorted); i++ {
		if sorted[i] <= 0 {
			// on time
			count++
		} else {
			// late
			break
		}
	}
	if count >= k {
		return "NO"
	}
	return "YES"
}

// already have merge sort function in this project but want to practice to memorize how it's done
func mSort(arr []int32) []int32 {
	// devide till reach only one element and return
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mSort(arr[:mid])
	right := mSort(arr[mid:])

	return mergeSwap(left, right)
}

func mergeSwap(left, right []int32) []int32 {
	result := make([]int32, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
