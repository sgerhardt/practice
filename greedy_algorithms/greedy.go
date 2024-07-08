package greedy_algorithms

//Given an array of integers, find the highest product you can get from three of the integers.
//The input arrayOfInts will always have at least three integers.

// TODO the below doesn't work with negatives
// There are at least two great answers:
//
// Keep track of the highest2 and lowest2 (most negative) numbers. If the current number times some combination of those is higher than the current highestProductOf3, we have a new highestProductOf3!
// Keep track of the highestProductOf2 and lowestProductOf2 (could be a low negative number). If the current number times one of those is higher than the current highestProductOf3, we have a new highestProductOf3!
func highestProductFromThreeInts(arrayOfInts []int) int {
	if len(arrayOfInts) < 3 {
		return 0
	}
	firstHighest := 0
	secondHighest := 0
	thirdHighest := 0
	for _, i := range arrayOfInts {
		if i > firstHighest {
			thirdHighest = secondHighest
			secondHighest = firstHighest
			firstHighest = i
		} else if i > secondHighest {
			thirdHighest = secondHighest
			secondHighest = i
		} else if i > thirdHighest {
			thirdHighest = i
		}
	}

	return firstHighest * secondHighest * thirdHighest
}
