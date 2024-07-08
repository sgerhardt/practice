package arrays_and_string_manipulation

import "sort"

//Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
//and return an array of the non-overlapping intervals that cover all the intervals in the input.
//
//
//
//Example 1:
//
//Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
//Output: [[1,6],[8,10],[15,18]]
//Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
//Example 2:
//
//Input: intervals = [[1,4],[4,5]]
//Output: [[1,5]]
//Explanation: Intervals [1,4] and [4,5] are considered overlapping.
//

func mergeIntervals(intervals [][]int) [][]int {
	if intervals == nil {
		return nil
	}

	var intervalRanges intervalRange
	intervalRanges = intervals
	sort.Sort(intervalRanges)

	lowestVal := 0
	highestVal := 0
	var merged [][]int
	for i := 0; i < len(intervalRanges); i++ {
		if i == 0 {
			lowestVal = intervalRanges[0][0]
			highestVal = intervalRanges[0][1]

			// nothing to merge on first iteration
			merged = append(merged, intervalRanges[0])
		} else if intervalRanges[i][0] > highestVal {
			// nothing to merge
			lowestVal = intervalRanges[i][0]
			highestVal = intervalRanges[i][1]

			merged = append(merged, intervalRanges[i])
		} else if intervalRanges[i][0] >= lowestVal {
			// adjust the previous entry to have a new high value
			if highestVal > intervalRanges[i][1] {
				intervalRanges[i-1][1] = highestVal
			} else {
				merged[len(merged)-1][1] = intervalRanges[i][1]
				highestVal = intervalRanges[i][1]
			}

		}
	}

	return merged
}

type intervalRange [][]int

func (rng intervalRange) Len() int { return len(rng) }

func (rng intervalRange) Less(i, j int) bool { return rng[i][0] < rng[j][0] }

func (rng intervalRange) Swap(i, j int) { rng[i], rng[j] = rng[j], rng[i] }
