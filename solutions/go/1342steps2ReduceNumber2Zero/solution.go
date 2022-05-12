package steps2ReduceNumber2Zero

// URL 		  - https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
// Difficulty - Easy

/*
 * Given an integer num, return the number of steps to reduce it to zero.
 *
 * In one step, if the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.
 */

func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	}
	steps := 0
	for num != 0 {
		if num&1 == 0 {
			steps += 1
		} else {
			steps += 2
		}
		num >>= 1
	}

	return steps - 1
}
