package roman_to_int

// URL 		  - https://leetcode.com/problems/roman-to-integer/
// Difficulty - Easy

func romanToInt(input string) int {
	result := 0
	romanNumbers := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for p := range input {
		if p+1 == len(input) {
			result += romanNumbers[input[p]]
		} else if romanNumbers[input[p]] >= romanNumbers[input[p+1]] {
			result += romanNumbers[input[p]]
		} else {
			result -= romanNumbers[input[p]]
		}
	}
	return result
}
