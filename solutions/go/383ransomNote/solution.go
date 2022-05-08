package ransomNote

// URL 		  - https://leetcode.com/problems/ransom-note/
// Difficulty - Easy

/*
 * Given two strings ransomNote and magazine, return true if ransomNote can be constructed from magazine and false otherwise.
 * Each letter in magazine can only be used once in ransomNote.
 */

func canConstruct(ransomNote string, magazine string) bool {
	ransomLetters := make(map[byte]int)
	magazineLetters := make(map[byte]int)

	for i := range ransomNote {
		ransomLetters[ransomNote[i]] += 1
	}

	for i := range magazine {
		magazineLetters[magazine[i]] += 1
	}

	for k, v1 := range ransomLetters {
		if v2, ok := magazineLetters[k]; v1 > v2 || !ok {
			return false
		}
	}

	return true
}
