package greedy_algorithms

//You have an array of integers, and for each index you want to find the product of every integer except the integer at that index.
//Write a method getProductsOfAllIntsExceptAtIndexBruteForce() that takes an array of integers and returns an array of the products.
//For example, given:
//
//  [1, 7, 3, 4]
//
//your method would return:
//
//  [84, 12, 28, 21]
//by calculating:
//
//  [7 * 3 * 4,  1 * 3 * 4,  1 * 7 * 4,  1 * 7 * 3]
//Here's the catch: You can't use division in your solution!

func getProductsOfAllIntsExceptAtIndexBruteForce(ints []int) []int {
	sumProducts := make([]int, len(ints))
	for i := 0; i < len(ints); i++ {
		sum := 1
		for j := 0; j < len(ints); j++ {
			if i == j {
				continue
			}
			sum *= ints[j]
		}
		sumProducts[i] = sum
	}
	return sumProducts
}

func getProductsOfAllIntsExceptAtIndex(ints []int) []int {
	// for each integer, we find the product of all the integers
	// before it, storing the total product so far each time
	productsOfAllIntsExceptAtIndex := make([]int, len(ints))
	productSoFar := 1
	for i := 0; i < len(ints); i++ {
		productsOfAllIntsExceptAtIndex[i] = productSoFar
		productSoFar *= ints[i]
	}

	// for each integer, we find the product of all the integers
	// after it. since each index in products already has the
	// product of all the integers before it, now we're storing
	// the total product of all other integers
	productSoFar = 1
	for i := len(ints) - 1; i >= 0; i-- {
		productsOfAllIntsExceptAtIndex[i] *= productSoFar
		productSoFar *= ints[i]
	}

	return productsOfAllIntsExceptAtIndex
}
