package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortScores(t *testing.T) {
	//  int[] unsortedScores = {37, 89, 41, 65, 91, 53};
	//final int HIGHEST_POSSIBLE_SCORE = 100;
	//
	//int[] sortedScores = sortScores(unsortedScores, HIGHEST_POSSIBLE_SCORE);
	// sortedScores: [91, 89, 65, 53, 41, 37]
	unsortedScores := []int{37, 89, 41, 65, 91, 53}
	sortedScoresExpected := []int{91, 89, 65, 53, 41, 37}
	assert.Equal(t, sortedScoresExpected, sortScores(unsortedScores, 100))

	unsortedScores = []int{37, 89, 41, 89, 91, 53}
	sortedScoresExpected = []int{91, 89, 89, 53, 41, 37}
	assert.Equal(t, sortedScoresExpected, sortScores(unsortedScores, 100))
}
