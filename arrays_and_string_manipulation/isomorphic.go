package arrays_and_string_manipulation

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Use slices to store the last seen positions of characters.
	// Assuming ASCII characters, we can use an array of size 256.
	var lastSeenS [256]int
	var lastSeenT [256]int

	for i := 0; i < len(s); i++ {
		charS := s[i]
		charT := t[i]

		// Check the last seen positions.
		if lastSeenS[charS] != lastSeenT[charT] {
			return false
		}

		// Update the last seen positions to the current index + 1
		// (index + 1 is used to differentiate from the default value 0).
		lastSeenS[charS] = i + 1
		lastSeenT[charT] = i + 1
	}

	return true
}
