package leetCode

func lengthOfLongestSubstring(s string) int {
	n := int32(len(s))
	ans := int32(0)
	charInts := make([]int32, len(s))
	for i, c := range s {
		charInts[i] = c
	}

	charToIndex := map[string]int32{}
	// using a "sliding window" which is two indexes keeping track of where we are in the string.
	for j, i := int32(0), int32(0); j < n; j++ {
		if index, ok := charToIndex[string(charInts[j])]; ok {
			// We've come across a duplicate character, so we bump our lower index
			i = getMax(index, i)
		}
		//keep track of our longest string
		ans = getMax(ans, j-i+1)
		charToIndex[string(charInts[j])] = j + 1
	}
	return int(ans)
}

func getMax(i, j int32) int32 {
	if i > j {
		return i
	} else {
		return j
	}
}
