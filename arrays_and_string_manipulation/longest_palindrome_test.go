package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	//assert.Equal(t, 3, longestPalindrome("aab")) // aba
	//assert.Equal(t, 2, longestPalindrome("bb"))
	//assert.Equal(t, 1, longestPalindrome("abc"))      // a, b or c
	assert.Equal(t, 7, longestPalindrome("abccccdd")) // dccaccd
}
