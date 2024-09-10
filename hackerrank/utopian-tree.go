package hackerrank

func UtopianTree(n int32) int32 {
	// Write your code here
	var i int32 = 1
	var height int32 = 1
	for i <= n {
		if i%2 == 0 {
			height += 1
		} else {
			height += height
		}
		i++
	}
	return height
}
