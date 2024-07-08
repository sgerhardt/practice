package arrays_and_string_manipulation

import (
	"strconv"
	"strings"
)

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
//
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
//
// Example 1:
//
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Example 2:
//
// Input: strs = [""]
// Output: [[""]]
// Example 3:
//
// Input: strs = ["a"]
// Output: [["a"]]
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	//anagramMap is a map of ascii frequency array to a list of strings
	// e.g. tea and eat would have a key of a1e1t1
	anagramMap := map[string][]string{}

	// make a frequency map of charToOccurrence
	for _, s := range strs {
		// chars for each char in alphabet
		frequencyArr := make([]int, 26)
		for _, c := range s {
			// subtract the char from 'a' ascii value, such that a starts at 0...
			frequencyArr[c-'a']++
		}

		// construct key from ascii values
		sb := strings.Builder{}
		for idx, value := range frequencyArr {
			if value == 0 {
				continue
			}
			char := string(rune(idx + 'a'))
			sb.WriteString(char + strconv.Itoa(value))
		}

		res, exists := anagramMap[sb.String()]
		if exists {
			anagramMap[sb.String()] = append(res, s)
		} else {
			anagramMap[sb.String()] = []string{s}
		}
	}

	var result [][]string
	for _, words := range anagramMap {
		result = append(result, words)
	}
	return result
}
