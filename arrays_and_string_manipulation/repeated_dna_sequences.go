package arrays_and_string_manipulation

// The DNA sequence is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T'.
//
// For example, "ACGAATTCCG" is a DNA sequence.
// When studying DNA, it is useful to identify repeated sequences within the DNA.
//
// Given a string s that represents a DNA sequence, return all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule. You may return the answer in any order.
//
// Example 1:
//
// Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
// Output: ["AAAAACCCCC","CCCCCAAAAA"]
// Example 2:
//
// Input: s = "AAAAAAAAAAAAA"
// Output: ["AAAAAAAAAA"]
func findRepeatedDnaSequences(s string) []string {
	sequenceSet := map[string]struct{}{}
	resultSet := map[string]struct{}{}

	for i := 0; i+10 <= len(s); i++ {
		sequence := s[i : i+10]
		_, exists := sequenceSet[sequence]
		if exists {
			resultSet[sequence] = struct{}{}
		} else {
			sequenceSet[sequence] = struct{}{}
		}
	}

	var results []string
	for result := range resultSet {
		results = append(results, result)
	}
	return results
}
