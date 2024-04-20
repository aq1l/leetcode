package validParentheses

// URL 		  - https://leetcode.com/problems/valid-parentheses"
// Difficulty - Easy

/*
 * Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *     - Open brackets must be closed by the same type of brackets.
 *     - Open brackets must be closed in the correct order.
 *     - Every close bracket has a corresponding open bracket of the same type.
 */

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else if len(stack) == 0 {
			return false
		} else if s[i] == ')' && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else if s[i] == '}' && stack[len(stack)-1] == '{' {
			stack = stack[:len(stack)-1]
		} else if s[i] == ']' && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
