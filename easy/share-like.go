package easy

func viralAdvertising(n int32) int32 {
	// Write your code here
	var share int32 = 5
	var floor int32 = share / 2
	var liked int32 = floor
	var cumulative int32 = liked
	// start at 1 and continue grow to 3
	var shareOther int32 = 3
	for i := int32(1); i < n; i++ {
		totalshare := liked * shareOther
		liked = totalshare / 2
		cumulative += liked
	}
	return cumulative
}
