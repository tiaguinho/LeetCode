/**
 * Link to the problem
 * https://leetcode.com/problems/remove-outermost-parentheses/
 */
package leetcode

func removeOuterParentheses(S string) string {
	outer := ""
	opened := 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			if opened > 0 {
				outer += string(S[i])
			}
			opened++
		} else if S[i] == ')' {
			if opened > 1 {
				outer += string(S[i])
			}
			opened--
		}
	}

	return outer
}
