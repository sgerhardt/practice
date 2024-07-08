package arrays_and_string_manipulation

import (
	"sort"
)

//Your company built an in-house calendar tool called HiCal. You want to add a feature to see the times in a day when everyone is available.
//
//To do this, you’ll need to know when any team is having a meeting. In HiCal, a meeting is stored as a tuple ↴ of integers (start_time, end_time).
//These integers represent the number of 30-minute blocks past 9:00am.
//
//For example:
//
// (2, 3)  # Meeting from 10:00 – 10:30 am
//(6, 9)  # Meeting from 12:00 – 1:30 pm
//
//Python 3.6
//Write a function merge_ranges() that takes a list of multiple meeting time ranges and returns a list of condensed ranges.
//
//For example, given:
//
//  [(0, 1), (3, 5), (4, 8), (10, 12), (9, 10)]
//
//Python 3.6
//your function would return:
//
//  [(0, 1), (3, 8), (9, 12)]
//
//Python 3.6
//Do not assume the meetings are in order. The meeting times are coming from multiple teams.
//
//Write a solution that's efficient even when we can't put a nice upper bound on the numbers representing our time ranges.
//Here we've simplified our times down to the number of 30-minute slots past 9:00 am.
//But we want the function to work even for very large numbers, like Unix timestamps.
//In any case, the spirit of the challenge is to merge meetings where start_time and end_time don't have an upper bound.

type meetingTimeframe struct {
	start int
	end   int
}

type meetingTimeframes []meetingTimeframe

func mergeRanges(times []meetingTimeframe) []meetingTimeframe {
	earliestStartTime := 0
	latestEndTime := 0
	latestTime := 0

	sort.Sort(meetingTimeframes(times))
	mergedRanges := []meetingTimeframe{}
	for i, t := range times {
		if t.end > latestTime {
			latestTime = t.end
		}

		if i == 0 {
			earliestStartTime = t.start
			latestEndTime = t.end

			// nothing to merge on first iteration...
			mergedRanges = append(mergedRanges, t)
			continue
		} else if t.start > latestEndTime {
			// nothing to merge
			earliestStartTime = t.start
			latestEndTime = t.end
			mergedRanges = append(mergedRanges, t)
			continue
		} else if t.start >= earliestStartTime {
			// adjust end time of previous entry to use new entry
			if t.end > mergedRanges[len(mergedRanges)-1].end {
				mergedRanges[len(mergedRanges)-1].end = t.end
			}
		}
	}

	return mergedRanges
}

func sortRanges(times []meetingTimeframe) []meetingTimeframe {
	sort.Sort(meetingTimeframes(times))
	return times
}

func (times meetingTimeframes) Len() int { return len(times) }

func (times meetingTimeframes) Less(i, j int) bool { return times[i].start < times[j].start }

func (times meetingTimeframes) Swap(i, j int) { times[i], times[j] = times[j], times[i] }
