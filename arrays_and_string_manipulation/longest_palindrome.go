package arrays_and_string_manipulation

// Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.
//
// Letters are case sensitive, for example, "Aa" is not considered a palindrome here.
//
// Example 1:
//
// Input: s = "abccccdd"
// Output: 7
// Explanation: One longest palindrome that can be built is "dccaccd", whose length is 7.
// Example 2:
//
// Input: s = "a"
// Output: 1
// Explanation: The longest palindrome that can be built is "a", whose length is 1.
func longestPalindrome(s string) int {
	if len(s) == 1 {
		return 1
	}

	// Use a set to keep track of characters we've encountered. If we've seen it before, delete it.
	// At the end,
	charSet := map[string]struct{}{}
	for _, c := range s {
		_, exists := charSet[string(c)]
		if exists {
			delete(charSet, string(c))
		} else {
			charSet[string(c)] = struct{}{}
		}
	}

	if len(charSet) != 0 {
		// odd length, e.g. abccccdd (dccaccd) returns 7. Length of string is 8, charSet is 2 (a and b), so 8 - 2 + 1 = 7.
		// add 1 since palindrome can contain 1 odd count character
		return len(s) - len(charSet) + 1
	}
	return len(s)
}
