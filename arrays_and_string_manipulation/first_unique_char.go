package arrays_and_string_manipulation

import "math"

// Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.
//
// Example 1:
//
// Input: s = "leetcode"
// Output: 0
// Example 2:
//
// Input: s = "loveleetcode"
// Output: 2
// Example 3:
//
// Input: s = "aabb"
// Output: -1
func firstUniqChar(s string) int {
	charIndex := map[string]int{}

	for idx, char := range s {
		val := string(char)
		_, exists := charIndex[val]
		if exists {
			charIndex[val] = -1
		} else {
			charIndex[val] = idx
		}
	}

	firstValidIdx := math.MaxInt
	for _, val := range charIndex {
		if val >= 0 && val < firstValidIdx {
			firstValidIdx = val
		}
	}
	if firstValidIdx == math.MaxInt {
		return -1
	}

	return firstValidIdx
}

// another idea...
//Get the frequency of each character.
//Get the first character that has a frequency of one.
//Actually the code below passes all the cases. However, according to @xietao0221, we could change the size of the frequency array to 256 to store other kinds of characters. Thanks for all the other comments and suggestions. Fight on!
//
//public class Solution {
//    public int firstUniqChar(String s) {
//        int freq [] = new int[26];
//        for(int i = 0; i < s.length(); i ++)
//            freq [s.charAt(i) - 'a'] ++;
//        for(int i = 0; i < s.length(); i ++)
//            if(freq [s.charAt(i) - 'a'] == 1)
//                return i;
//        return -1;
//    }
//}
