package hackerrank

func saveThePrisoner(n int32, m int32, s int32) int32 {
	// Write your code here
	// n is a number of seat
	// m number of candies
	// s is starting chair
	seatNo := (s + m - 1) % n
	if seatNo == 0 {
		return n
	}
	return seatNo

	// for i:=int32(1) ; i <= m ; i++ {
	//     if i == m {
	//         return s
	//     }

	//     if s == n {
	//         s = 1
	//         continue
	//     }
	//     s++
	// }

	// return chair number which mean s that count to
	// return s
}
