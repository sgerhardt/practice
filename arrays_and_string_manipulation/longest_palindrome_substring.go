package arrays_and_string_manipulation

// Given a string s, return the longest palindromic substring in s.
//
// Example 1:
//
// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.
// Example 2:
//
// Input: s = "cbbd"
// Output: "bb"
func longestPalindromeSubstring(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	res := ""

	for i := 0; i < len(s); i++ {
		// odd case like "aba"
		tmp := helper(s, i, i)
		if len(tmp) > len(res) {
			res = tmp
		}

		// even case like "abba"
		tmp = helper(s, i, i+1)
		if len(tmp) > len(res) {
			res = tmp
		}
	}

	return res
}

func helper(str string, left, right int) string {
	for left >= 0 && right < len(str) && str[left] == str[right] {
		left--
		right++
	}
	return str[left+1 : right]
}
