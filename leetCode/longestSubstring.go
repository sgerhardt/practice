package leetCode

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[rune]int)
	// start is where we start counting from in the string
	start, maxLength := 0, 0
	for i, char := range s {
		if prevIndex, found := charIndex[char]; found && prevIndex >= start {
			// if we found a duplicate character, we move the left index that we search from
			start = prevIndex + 1
		}
		charIndex[char] = i
		// get the current length of the string, which is the current index and set it to max if necessary
		currentLength := i - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}
