package validpalindrome

// URL 		  - https://leetcode.com/problems/valid-palindrome
// Difficulty - Easy

/*
 * A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,
 * it reads the same forward and backward. Alphanumeric characters include letters and numbers.
 *
 * Given a string s, return true if it is a palindrome, or false otherwise.
 */

func isPalindrome(s string) bool {
	r := []rune(s)
	i, j := 0, len(r)-1
	for i < j {
		for i < j && !((r[i] >= 'a' && r[i] <= 'z') || (r[i] >= 'A' && r[i] <= 'Z') || (r[i] >= '0' && r[i] <= '9')) {
			i++
		}

		for i < j && !((r[j] >= 'a' && r[j] <= 'z') || (r[j] >= 'A' && r[j] <= 'Z') || (r[j] >= '0' && r[j] <= '9')) {
			j--
		}

		if r[i] >= 'A' && r[i] <= 'Z' {
			r[i] = r[i] + 32
		}

		if r[j] >= 'A' && r[j] <= 'Z' {
			r[j] = r[j] + 32
		}

		if r[i] != r[j] {
			return false
		}

		i++
		j--
	}

	return true
}
