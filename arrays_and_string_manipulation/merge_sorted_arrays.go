package arrays_and_string_manipulation

import "sort"

// Each order is represented by an "order id" (an integer).
//
// We have our lists of orders sorted numerically already, in lists. Write a function to merge our lists of orders into one sorted list.
//
// For example:
// my_list     = [3, 4, 6, 10, 11, 15]
// alices_list = [1, 5, 8, 12, 14, 19]
//
// # Prints [1, 3, 4, 5, 6, 8, 10, 11, 12, 14, 15, 19]
// Uses nlogn with built in sort.Sort. However, this _can_ be done in n time and space!
type intList []int

func mergeSortedNlogN(firstList intList, secondList intList) []int {
	thirdList := append(firstList, secondList...)
	sort.Sort(thirdList)
	return thirdList
}

func (list intList) Len() int { return len(list) }

func (list intList) Less(i, j int) bool { return list[i] < list[j] }

func (list intList) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

func mergeSortedN(firstList intList, secondList intList) []int {
	mergedList := make(intList, len(firstList)+len(secondList))

	mergedListSize := len(firstList) + len(secondList)
	firstListIndex := 0
	secondListIndex := 0
	mergedListIndex := 0

	for mergedListIndex < mergedListSize {

		firstListMerged := false
		secondListMerged := false
		var firstListUnmergedValue int
		var secondListUnmergedValue int
		if firstListIndex == len(firstList) {
			// first all merged
			firstListMerged = true
		}
		if secondListIndex == len(secondList) {
			// second all merged
			secondListMerged = true
		}
		if firstListMerged && secondListMerged {
			break
		}

		if !firstListMerged {
			firstListUnmergedValue = firstList[firstListIndex]
		} else {
			mergedList[mergedListIndex] = secondList[secondListIndex]
			secondListIndex += 1
			continue
		}

		if !secondListMerged {
			secondListUnmergedValue = secondList[secondListIndex]
		} else {
			mergedList[mergedListIndex] = firstList[firstListIndex]
			firstListIndex += 1
		}

		if firstListUnmergedValue < secondListUnmergedValue {
			mergedList[mergedListIndex] = firstListUnmergedValue
			firstListIndex += 1
		} else {
			mergedList[mergedListIndex] = secondListUnmergedValue
			secondListIndex += 1
		}
		mergedListIndex++
	}
	return mergedList
}
