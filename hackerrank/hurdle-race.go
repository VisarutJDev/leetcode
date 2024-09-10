package hackerrank

func HurdleRace(k int32, height []int32) int32 {
	// Write your code here

	// now you get the max value from last value
	newHeight := MergeSort(height)
	max := newHeight[len(newHeight)-1]

	moreDoseRequired := max - k
	if moreDoseRequired < 0 {
		return 0
	}
	return moreDoseRequired
}

func MergeSort(arr []int32) []int32 {
	// divide till it last element
	if len(arr) <= 1 {
		return arr
	}

	// divide by half
	midTodivide := len(arr) / 2
	left := MergeSort(arr[:midTodivide])
	right := MergeSort(arr[midTodivide:])

	// Merge sort algorithm
	// just swap element
	return merge(left, right)
}

func merge(left, right []int32) []int32 {
	// init result that merge left and right
	result := make([]int32, 0, len(left)+len(right))
	// index to iterate left and right
	i, j := 0, 0

	// loop while index of both left and right still in range
	for i < len(left) && j < len(right) {
		// find less value and put it in result
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// because i can have more to go when j is already more than len(right)
	// just like j can have more to go when i is already more than len(left)

	// remaining element
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
