package dynamic_programming_and_recursion

import "fmt"

//Your quirky boss collects rare, old coins...
//
//They found out you're a programmer and asked you to solve something they've been wondering for a long time.
//
//Write a function that, given:
//
//an amount of money
//a list of coin denominations
//computes the number of ways to make the amount of money with coins of the available denominations.
//
//Example: for amount=4 (4¢) and denominations=[1,2,3] (1¢, 2¢ and 3¢), your program would output 4—the number of ways to make 4¢ with those denominations:
//
//1¢, 1¢, 1¢, 1¢
//1¢, 1¢, 2¢
//1¢, 3¢
//2¢, 2¢

// changePossibilitiesTopDownRecursive returns the number of ways to make the amount left. Could improve with memoization
func changePossibilitiesTopDownRecursive(amountLeft int, denominations []int, currentIndex int) int {
	if amountLeft == 0 {
		return 1
	}

	// we overshot amount, used too many coins
	if amountLeft < 0 {
		return 0
	}

	if currentIndex == len(denominations) {
		return 0
	}

	fmt.Printf("Checking ways to make :%+v with denomination:%+v\n", amountLeft, denominations[currentIndex:])
	currentCoin := denominations[currentIndex]

	// See how many possibilities we can get
	// for each number of times to use currentCoin
	numPossibilities := 0
	for amountLeft >= 0 {
		numPossibilities += changePossibilitiesTopDownRecursive(amountLeft, denominations, currentIndex+1)
		amountLeft -= currentCoin
	}
	return numPossibilities
}

// With the bottom up approach, observe the pattern when starting with one denomination and going to more...

// We can start by making an array waysOfDoingNCents, where the index is the amount and
// the value at each index is the number of ways of getting that amount.

// To further simplify the problem, we can work with only the first coin in denominations, then add in the second coin, then the third, etc.
//  int[] waysOfDoingNCents1 = new int[] {
//    1,  // 0c:  no coins
//    1,  // 1c:  1 1c coin
//    1,  // 2c:  2 1c coins
//    1,  // 3c:  3 1c coins
//    1,  // 4c:  4 1c coins
//    1,  // 5c:  5 1c coins
//};
//  int[] waysOfDoingNCents1And2 = new int[] {
//    1,      // 0c:  no change
//    1,      // 1c:  no change
//    1 + 1,  // 2c:  new [(2)]
//    1 + 1,  // 3c:  new [(2, 1)]
//    1 + 2,  // 4c:  new [(2, 1, 1), (2,2)]
//    1 + 2,  // 5c:  new [(2, 1, 1, 1), (2, 2, 1)]

//	 int[] waysOfDoingNCents1And2And3 = new int[] {
//	   1,      // 0c:  no change
//	   1,      // 1c:  no change
//	   2,      // 2c:  no change
//	   2 + 1,  // 3c:  new [(3)]
//	   3 + 1,  // 4c:  new [(3, 1)]
//	   3 + 2,  // 5c:  new [(3, 2), (3, 1, 1)]
//	};
//
// Notice how we increment ways of doing N cents starting at the index of whatever the N cent is, e.g. here we start incrementing possibilities at index=3
func changePossibilitiesBottomUpIterative(amount int, denominations []int) int {
	waysOfDoingNCents := make([]int, amount+1) // array of zeros from 0..amount
	// our base case - there is _one_ way to get 0 cents...
	waysOfDoingNCents[0] = 1

	for _, coin := range denominations {
		// start at the coin, then increment until we reach the amount
		for higherAmount := coin; higherAmount <= amount; higherAmount++ {
			higherAmountRemainder := higherAmount - coin
			waysOfDoingNCents[higherAmount] += waysOfDoingNCents[higherAmountRemainder]
		}
	}
	return waysOfDoingNCents[amount]
}
