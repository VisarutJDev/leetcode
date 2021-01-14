package easy

// MaximumWealth question
// You are given an m x n integer grid accounts where accounts[i][j] is the amount of money the i​​​​​​​​​​​th​​​​ customer has in the j​​​​​​​​​​​th​​​​ bank. Return the wealth that the richest customer has.
// A customer's wealth is the amount of money they have in all their bank accounts. The richest customer is the customer that has the maximum wealth.
func MaximumWealth(accounts [][]int) int {
	max := 0
	for _, acc := range accounts {
		sum := 0
		for _, a := range acc {
			sum = sum + a
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
