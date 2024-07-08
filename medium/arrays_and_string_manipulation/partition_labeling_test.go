package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartition(t *testing.T) {
	//Input: s = "ababcbacadefegdehijhklij"
	//Output: [9,7,8]
	//Explanation:
	//The partition is "ababcbaca", "defegde", "hijhklij".
	//This is a partition so that each letter appears in at most one part.
	//A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits s into less parts.
	assert.Equal(t, []int{9, 7, 8}, partitionLabels("ababcbacadefegdehijhklij"))
}
