package arrays_and_string_manipulation

//You are given a string s. We want to partition the string into as many parts as possible so that each letter appears in at most one part.
//
//Note that the partition is done so that after concatenating all the parts in order, the resultant string should be s.
//
//Return a list of integers representing the size of these parts.
//Input: s = "ababcbacadefegdehijhklij"
//Output: [9,7,8]
//Explanation:
//The partition is "ababcbaca", "defegde", "hijhklij".
//This is a partition so that each letter appears in at most one part.
//A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits s into less parts.

// Let's try to repeatedly choose the smallest left-justified partition.
// Consider the first label, say it's 'a'. The first partition must include it, and also the last occurrence of 'a'.
// However, between those two occurrences of 'a', there could be other labels that make the minimum size of this partition bigger.
// For example, in "abccaddbeffe", the minimum first partition is "abccaddb".
// This gives us the idea for the algorithm: For each letter encountered, process the last occurrence of that letter, extending the current partition [anchor, j] appropriately.
func partitionLabels(s string) []int {
	// Find the last occurrence for each character
	var lastOccurrence = make([]int, 26)
	for i := 0; i < len(s); i++ {
		// this gets us a zero-based index for the alphabet. ascii lowercase a starts at 97, so if the character is b (98), it will be in index 1.
		alphaIndex := s[i] - 'a'
		lastOccurrence[alphaIndex] = i
	}

	j := 0
	// the anchor is the starting position of the partition we are currently processing.
	anchor := 0
	var result []int
	for i := 0; i < len(s); i++ {
		// check to see if the character would expand the partition...
		j = max(j, lastOccurrence[s[i]-'a'])
		if i == j {
			result = append(result, i-anchor+1)
			anchor = i + 1
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
