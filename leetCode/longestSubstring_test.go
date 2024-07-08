package leetCode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	assert.Equal(t, 3, lengthOfLongestSubstring("dvdf"))
	assert.Equal(t, 6, lengthOfLongestSubstring("asjrgapa"))
	assert.Equal(t, 4, lengthOfLongestSubstring("jbpnbwwd"))

}
