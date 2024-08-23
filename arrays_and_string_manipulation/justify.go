package arrays_and_string_manipulation

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	var result []string
	var currentLine []string
	currentLineLength := 0

	for _, word := range words {
		// Check if adding this word would exceed the maxWidth
		if currentLineLength+len(word)+len(currentLine) > maxWidth {
			result = append(result, justify(currentLine, currentLineLength, maxWidth))
			currentLine = nil
			currentLineLength = 0
		}
		// Add the word to the current line
		currentLine = append(currentLine, word)
		currentLineLength += len(word)
	}

	// Handle the last line: left-justified
	lastLine := strings.Join(currentLine, " ")
	lastLine += strings.Repeat(" ", maxWidth-len(lastLine))
	result = append(result, lastLine)

	return result
}

func justify(words []string, wordsLength, maxWidth int) string {
	if len(words) == 1 {
		// Only one word, left-justify it
		return words[0] + strings.Repeat(" ", maxWidth-wordsLength)
	}

	// Calculate the number of spaces between words
	totalSpaces := maxWidth - wordsLength
	spaceWidth := totalSpaces / (len(words) - 1)
	extraSpaces := totalSpaces % (len(words) - 1)

	var result string
	for i, word := range words {
		result += word
		if i < len(words)-1 {
			result += strings.Repeat(" ", spaceWidth)
			if extraSpaces > 0 {
				result += " "
				extraSpaces--
			}
		}
	}

	return result
}
