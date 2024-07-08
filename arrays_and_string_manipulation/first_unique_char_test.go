package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstUniqueChar(t *testing.T) {
	assert.Equal(t, 0, firstUniqChar("leetcode"))
	assert.Equal(t, 2, firstUniqChar("loveleetcode"))
	assert.Equal(t, -1, firstUniqChar("aabb"))
}
