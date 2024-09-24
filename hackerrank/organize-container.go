package hackerrank

import "sort"

func organizingContainers(container [][]int32) string {
	// Write your code here
	// so the pattern is
	// Number of balls in each type should be equal to be possible
	// if not impossible
	// this problem is about diagonal matrix
	capacityOfContainer := make([]int32, len(container))
	sumOfBallsType := make([]int32, len(container))
	for i := 0; i < len(container); i++ {
		for j := 0; j < len(container[i]); j++ {
			capacityOfContainer[i] += container[i][j]
			sumOfBallsType[j] += container[i][j]
		}
	}
	sort.Slice(capacityOfContainer, func(i, j int) bool {
		return capacityOfContainer[i] < capacityOfContainer[j]
	})
	sort.Slice(sumOfBallsType, func(i, j int) bool {
		return sumOfBallsType[i] < sumOfBallsType[j]
	})
	for i := 0; i < len(capacityOfContainer); i++ {
		if sumOfBallsType[i] > capacityOfContainer[i] {
			return "Impossible"
		}
	}
	// for i := 0; i< len(capacityOfContainer); i++ {
	//      for j:= 0; j<len(sumOfBallsType); j++ {

	//          if sumOfBallsType[j] > capacityOfContainer[i] {
	//              fmt.Println(capacityOfContainer[i], sumOfBallsType[j])
	//              return "Impossible"
	//          }
	//      }
	// }

	// [1 4] => 5
	// [2 3] => 5
	//  3 7

	// [1 1] => 2 => [2 0]
	// [1 1] => 2 => [0 2]
	//  2 2

	// [1 3 1] => 5
	// [2 1 2] => 5
	// [3 3 3] => 9
	//  6 7 6

	// [0 2 1] => 3 => [1 2 0] => [1 2 0] => [2 1 0] => [3 0 0]
	// [1 1 1] => 3 => [1 1 1] => [2 1 0] => [1 2 0] => [0 3 0]
	// [2 0 0] => 2 => [1 0 1] => [0 0 2] => [0 0 2] => [0 0 2]
	//  3 3 2
	return "Possible"
}
