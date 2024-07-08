package arrays_and_string_manipulation

import "strings"

//Given a pattern and a string s, find if s follows the same pattern.
//
//Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.
//
//
//
//Example 1:
//
//Input: pattern = "abba", s = "dog cat cat dog"
//Output: true
//Example 2:
//
//Input: pattern = "abba", s = "dog cat cat fish"
//Output: false
//Example 3:
//
//Input: pattern = "aaaa", s = "dog cat cat dog"
//Output: false

func wordPattern(pattern string, s string) bool {
	charToWord := map[string]string{}

	// also need wordToChar, e.g. think about "abba", "dog dog dog dog".
	// In other words, if a word already exists for a mapped character, then we know the pattern is borked, since dog can't map to both
	// a AND b...
	wordToChar := map[string]string{}

	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	for i := 0; i < len(words); i++ {
		c := pattern[i : i+1]
		val, exists := charToWord[c]
		if exists {
			if val != words[i] {
				return false
			}
		} else {
			val, exists = wordToChar[words[i]]
			if exists {
				return false
			} else {
				charToWord[c] = words[i]
				wordToChar[words[i]] = c
			}
		}
	}
	return true
}
