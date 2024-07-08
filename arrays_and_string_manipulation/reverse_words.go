package arrays_and_string_manipulation

func reverseWordsInPlace(chars []byte) {
	//Note reverseWordsInPlace reverses IN PLACE, which means we modify the original data structure.
	// This is only to be done if there are performance limitations - typically we prefer "out of place" changes, e.g.
	// simply returning a new copy of the data rather than modify the original.
	// NOTE: THIS ONLY WORKS FOR SINGLE BYTE ASCII CHARACTERS!!
	allWords := [][]byte{}
	currentWord := []byte{}
	for i, c := range chars {
		if len(chars)-1 == i {
			// last word - add a space for it since it will become the first word
			currentWord = append(currentWord, c)
			currentWord = append(currentWord, byte(32))
			allWords = append(allWords, currentWord)
			currentWord = []byte{}
		}
		if string(c) == ` ` {
			// the first word becomes the last. Only add space if not first word
			if len(allWords) == 0 {
				allWords = append(allWords, currentWord)
			} else {
				currentWord = append(currentWord, c)
				allWords = append(allWords, currentWord)
			}
			currentWord = []byte{}
		} else {
			// keep building up existing word
			currentWord = append(currentWord, c)
		}
	}
	// reverse words
	index := 0
	for i := len(allWords); i > 0; i-- {
		word := allWords[i-1]
		for i, char := range word {
			chars[i+index] = char
		}
		index += len(word)
	}
}
