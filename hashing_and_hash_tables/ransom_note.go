package hashing_and_hash_tables

// Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
//
// Each letter in magazine can only be used once in ransomNote.
// Input: ransomNote = "a", magazine = "b"
// Output: false
// Example 2:
//
// Input: ransomNote = "aa", magazine = "ab"
// Output: false
// Example 3:
//
// Input: ransomNote = "aa", magazine = "aab"
// Output: true
func canConstructNaive(ransomNote string, magazine string) bool {
	megazineCharToCount := map[string]int{}
	ransomCharToCount := map[string]int{}

	for _, char := range magazine {
		megazineCharToCount[string(char)]++
	}
	for _, char := range ransomNote {
		ransomCharToCount[string(char)]++
	}

	for char, count := range ransomCharToCount {
		megaCount, exists := megazineCharToCount[char]
		if !exists {
			return false
		}
		if megaCount < count {
			return false
		}
	}
	return true
}

// a more space efficient route would be to make an array of chars equal to the alphabet, add to each index as we pass through the megazine,
// then do another iteration to subtract from those indexes for each character in the ransom note. if it's less than 0, return false
func canConstruct(ransomNote string, magazine string) bool {
	alphabet := make([]int, 26)
	for _, char := range magazine {
		alphabet[char-'a']++
	}

	for _, char := range ransomNote {
		alphabet[char-'a']--
		if alphabet[char-'a'] < 0 {
			return false
		}
	}
	return true
}
