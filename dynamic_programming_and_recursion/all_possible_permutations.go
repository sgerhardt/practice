package dynamic_programming_and_recursion

func getPermutations(input string) map[string]struct{} {
	// base case
	if len(input) <= 1 {
		return map[string]struct{}{input: struct{}{}}
	}

	allCharsExceptLast := input[:len(input)-1]
	lastChar := input[len(input)-1:]

	// recursive call: get all possible permutations for all chars except last
	permutationsOfAllCharsExceptLast := getPermutations(allCharsExceptLast)

	// put the last char in all possible positions for each of the above permutations
	permutations := map[string]struct{}{}
	for p := range permutationsOfAllCharsExceptLast {
		for i := 0; i <= len(allCharsExceptLast); i++ {
			permutation := p[:i] + lastChar + p[i:]
			permutations[permutation] = struct{}{}
		}
	}
	return permutations
}
