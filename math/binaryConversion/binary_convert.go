package binaryConversion

// Input: a = "1010", b = "1011"
// Output: "10101"
func addBinary(a string, b string) string {
	var longestSlice []rune
	var shortestSlice []rune
	if len(a) > len(b) {
		longestSlice = []rune(a)
		shortestSlice = []rune(b)
	} else {
		longestSlice = []rune(b)
		shortestSlice = []rune(a)
	}

	binaryResultStr := ""
	carry := false
	runeZero := []rune(`0`)[0]
	runeOne := []rune(`1`)[0]
	for len(longestSlice) > len(shortestSlice) {
		shortestSlice = append([]rune(`0`), shortestSlice...)
	}
	for i := len(longestSlice) - 1; i >= 0; i-- {
		if i <= len(shortestSlice)-1 {
			if shortestSlice[i] == runeZero && longestSlice[i] == runeZero {
				if carry {
					binaryResultStr = "1" + binaryResultStr
				} else {
					binaryResultStr = "0" + binaryResultStr
				}
				carry = false
			} else if shortestSlice[i] == runeOne && longestSlice[i] == runeZero || shortestSlice[i] == runeZero && longestSlice[i] == runeOne {
				if carry {
					binaryResultStr = "0" + binaryResultStr
					carry = true
				} else {
					binaryResultStr = "1" + binaryResultStr
					carry = false
				}
			} else if shortestSlice[i] == runeOne && longestSlice[i] == runeOne {
				if carry {
					binaryResultStr = "1" + binaryResultStr
				} else {
					binaryResultStr = "0" + binaryResultStr
				}
				carry = true
			}
			continue
		}
		if carry {
			if shortestSlice[i] == runeOne {
				binaryResultStr = "0" + binaryResultStr
				carry = true
			} else if shortestSlice[i] == runeZero {
				binaryResultStr = "1" + binaryResultStr
				carry = false
			}
		} else {
			binaryResultStr = string(longestSlice[i]) + binaryResultStr
			carry = false
		}
	}
	if carry {
		binaryResultStr = "1" + binaryResultStr
	}
	return binaryResultStrgo
}
