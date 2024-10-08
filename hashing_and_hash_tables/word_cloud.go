package hashing_and_hash_tables

import "strings"

//You want to build a word cloud,
//an infographic where the size of a word corresponds to how often it appears in the body of text.
//
//To do this, you'll need data. Write code that takes a long string and builds its word cloud data in a hash map,
//where the keys are words and the values are the number of times the words occurred.
//"After beating the eggs, Dana read the next step:"
//"Add milk and eggs, then add flour and sugar."
//What do we want to do with "After", "Dana", and "add"? In this example,
//your final hash map should include one "Add" or "add" with a value of 2.
//Make reasonable (not necessarily perfect) decisions about cases like "After" and "Dana".
//Assume the input will only contain words and standard punctuation.

func wordCloud(input string) map[string]int {
	wordToOccurrences := map[string]int{}

	splitStrings := strings.Split(input, " ")

	for _, str := range splitStrings {
		if strings.HasSuffix(str, ",") {
			str = str[:len(str)-1]
		} else if strings.HasSuffix(str, ":") {
			str = str[:len(str)-1]
		}
		wordToOccurrences[strings.ToLower(str)]++
	}
	return wordToOccurrences
}
