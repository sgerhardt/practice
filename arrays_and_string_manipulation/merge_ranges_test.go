package arrays_and_string_manipulation

import (
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestMergeRanges(t *testing.T) {
	//[(0, 1), (3, 5), (4, 8), (10, 12), (9, 10)]
	ranges := []meetingTimeframe{
		{
			start: 0,
			end:   1,
		},
		{
			start: 3,
			end:   5,
		},
		{
			start: 4,
			end:   8,
		},
		{
			start: 10,
			end:   12,
		},
		{
			start: 9,
			end:   10,
		},
	}
	merged := mergeRanges(ranges)

	expected := []meetingTimeframe{
		{
			start: 0,
			end:   1,
		},
		{
			start: 3,
			end:   8,
		},
		{
			start: 9,
			end:   12,
		},
	}
	assert.Equal(t, expected, merged)
}

func TestMultiMerge(t *testing.T) {
	//  [(1, 10), (2, 6), (3, 5), (7, 9)]
	ranges := []meetingTimeframe{
		{
			start: 1,
			end:   10,
		},
		{
			start: 2,
			end:   6,
		},
		{
			start: 3,
			end:   5,
		},
		{
			start: 7,
			end:   9,
		},
	}
	expected := []meetingTimeframe{
		{
			start: 1,
			end:   10,
		},
	}
	merged := mergeRanges(ranges)
	assert.Equal(t, expected, merged)
}

func TestOverlap(t *testing.T) {
	//[(1, 2), (2, 3)]
	ranges := []meetingTimeframe{
		{
			start: 1,
			end:   2,
		},
		{
			start: 2,
			end:   3,
		},
	}
	merged := mergeRanges(ranges)

	expected := []meetingTimeframe{
		{
			start: 1,
			end:   3,
		},
	}
	assert.Equal(t, expected, merged)
}

func TestSubsume(t *testing.T) {
	//[(1, 5), (2, 3)]
	ranges := []meetingTimeframe{
		{
			start: 1,
			end:   5,
		},
		{
			start: 2,
			end:   3,
		},
	}
	merged := mergeRanges(ranges)

	expected := []meetingTimeframe{
		{
			start: 1,
			end:   5,
		},
	}
	assert.Equal(t, expected, merged)
}

func TestSortRanges(t *testing.T) {
	ranges := []meetingTimeframe{
		{
			start: 0,
			end:   1,
		},
		{
			start: 3,
			end:   5,
		},
		{
			start: 4,
			end:   8,
		},
		{
			start: 10,
			end:   12,
		},
		{
			start: 9,
			end:   10,
		},
	}
	sortedRanges := sortRanges(ranges)

	expectedRanges := []meetingTimeframe{
		{
			start: 0,
			end:   1,
		},
		{
			start: 3,
			end:   5,
		},
		{
			start: 4,
			end:   8,
		},
		{
			start: 9,
			end:   10,
		},
		{
			start: 10,
			end:   12,
		},
	}
	assert.Equal(t, expectedRanges, sortedRanges)
}
