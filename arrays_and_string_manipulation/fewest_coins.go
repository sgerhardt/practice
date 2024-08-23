package arrays_and_string_manipulation

func fewestCoins(value int) []int {
	// e.g. 51 cents gives 2 quarters and a penny
	// loop over denominations desc, if we have a remainder go to next smaller denom
	denoms := []int{25, 10, 5, 1}

	denomToCoinCount := make([]int, 4)
	for idx, d := range denoms {
		total := value / d
		denomToCoinCount[idx] = total
		value = value % d
	}

	return denomToCoinCount
}
