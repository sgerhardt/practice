package hashing_and_hash_tables

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
//
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
//
// Example 1:
//
// Input: s = "anagram", t = "nagaram"
// Output: true
// Example 2:
//
// Input: s = "rat", t = "car"
// Output: false
func isAnagram(s string, t string) bool {
	alphabet := make([]int, 26)
	for _, char := range s {
		alphabet[char-'a']++
	}

	for _, char := range t {
		alphabet[char-'a']--
	}

	for _, val := range alphabet {
		if val != 0 {
			return false
		}
	}
	return true
}
