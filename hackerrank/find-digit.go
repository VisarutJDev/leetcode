package hackerrank

func findDigits(n int32) int32 {
	// Write your code here
	var mResult []int32
	tmp := n
	for n != 0 {
		d := n % 10
		if d == 0 {
			n = n / 10
			continue
		}
		if tmp%d == 0 {
			mResult = append(mResult, d)
		}
		n = n / 10
	}

	return int32(len(mResult))
}
