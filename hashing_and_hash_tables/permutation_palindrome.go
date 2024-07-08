package hashing_and_hash_tables

//Write an efficient function that checks whether any permutation ↴ of an input string is a palindrome. ↴
//
//You can assume the input string only contains lowercase letters.
//
//Examples:
//
//"civic" should return True
//"ivicc" should return True
//"civil" should return False
//"livci" should return False

func palindromePossible(input string) bool {
	if len(input) == 0 || len(input) == 1 {
		return true
	}

	charToOccurrence := map[string]int{}
	if len(input)%2 == 0 {
		// Even palindrome. All characters must show up the same number of times
		for _, c := range input {
			println(string(c))
			charToOccurrence[string(c)]++
		}
		for _, occurrence := range charToOccurrence {
			if occurrence%2 != 0 {
				return false
			}
		}
		return true
	} else {
		// Odd palindrome. All characters, except for one, must have an even number of occurrences.
		for _, c := range input {
			println(string(c))
			charToOccurrence[string(c)]++
		}
		oddCharOccurrences := 0
		for _, occurrence := range charToOccurrence {
			if oddCharOccurrences > 1 {
				return false
			}
			if occurrence%2 != 0 {
				oddCharOccurrences++
			}
		}

		if oddCharOccurrences == 1 {
			return true
		}
	}
	return false
}
