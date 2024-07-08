package arrays_and_string_manipulation

import "sort"

// Given an array of intervals intervals where intervals[i] = [starti, endi],
// return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.
//
// Example 1:
//
// Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
// Output: 1
// Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.
// Example 2:
//
// Input: intervals = [[1,2],[1,2],[1,2]]
// Output: 2
// Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.
// Example 3:
//
// Input: intervals = [[1,2],[2,3]]
// Output: 0
// Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
// Actually, the problem is the same as "Given a collection of intervals, find the maximum number of intervals that are non-overlapping."
// (the classic Greedy problem: Interval Scheduling)
//
// Sorting Interval.end in ascending order is O(nlogn), then traverse intervals array to get the maximum number of non-overlapping intervals is O(n). Total is O(nlogn).
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	var intervalRanges intervalRangeRemovals
	intervalRanges = intervals
	sort.Sort(intervalRanges)

	count := 0
	end := intervalRanges[0][1]
	for i := 1; i < len(intervalRanges); i++ {
		if intervalRanges[i][0] < end {
			count++
		} else {
			end = intervalRanges[i][1]
		}
	}
	return count
}

type intervalRangeRemovals [][]int

func (rng intervalRangeRemovals) Len() int { return len(rng) }

func (rng intervalRangeRemovals) Less(i, j int) bool {
	return rng[i][1] < rng[j][1]
}

func (rng intervalRangeRemovals) Swap(i, j int) { rng[i], rng[j] = rng[j], rng[i] }
