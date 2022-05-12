package richestCustomerWealth

// URL 		  - https://leetcode.com/problems/richest-customer-wealth/
// Difficulty - Easy

/*
 * You are given an m x n integer grid accounts where accounts[i][j] is the amount of money the i-th customer has in the j-th bank.
 * Return the wealth that the richest customer has.
 *
 * A customer's wealth is the amount of money they have in all their bank accounts. The richest customer is the customer that has the maximum wealth.
 */

func maximumWealth(accounts [][]int) int {
	answer := 0
	for i := range accounts {
		sum := 0
		for j := range accounts[i] {
			sum += accounts[i][j]
		}
		if sum > answer {
			answer = sum
		}
	}
	return answer
}
