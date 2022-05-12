package kWeakestRowsInMatrix

// URL 		  - https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/
// Difficulty - Easy

/*
 * You are given an m x n binary matrix mat of 1's (representing soldiers) and 0's (representing civilians). The soldiers are positioned in front of the civilians. That is, all the 1's will appear to the left of all the 0's in each row.
 *
 * A row i is weaker than a row j if one of the following is true:
 *
 *     The number of soldiers in row i is less than the number of soldiers in row j.
 *     Both rows have the same number of soldiers and i < j.
 *
 * Return the indices of the k weakest rows in the matrix ordered from weakest to strongest.
 */

func kWeakestRows(mat [][]int, k int) []int {
	strengthSlice := make([]int, len(mat))

	for i := range mat {
		sum := 0
		for j := range mat[i] {
			sum += mat[i][j]
		}

		strengthSlice[i] = sum
	}

	var answer []int
	for i := 0; i < k; i++ {
		index := 0
		min := 101
		for j := range strengthSlice {
			if strengthSlice[j] < min {
				min = strengthSlice[j]
				index = j
			}

		}
		strengthSlice[index] = 100000
		answer = append(answer, index)
	}

	return answer
}
