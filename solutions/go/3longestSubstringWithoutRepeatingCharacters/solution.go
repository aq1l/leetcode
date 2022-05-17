package longestSubstringWithoutRepeatingCharacters

// URL 		  - https://leetcode.com/problems/longest-substring-without-repeating-characters/
// Difficulty - Medium

/*
 * Given a string s, find the length of the longest substring without repeating characters.
 */

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxLength := 0
	i := 0
	st := 0
	letters := make(map[byte]int)
	for i = range s {
		if _, ok := letters[s[i]]; ok {
			if letters[s[i]] >= st {
				if maxLength < i-st {
					maxLength = i - st
				}

				// move window
				st = letters[s[i]] + 1
			}

		}
		letters[s[i]] = i
	}

	if maxLength <= i-st {
		maxLength = i - st + 1
	}

	return maxLength
}
