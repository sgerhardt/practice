package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	assert.False(t, isPalindrome("ab"))
	assert.True(t, isPalindrome("abba"))
	assert.True(t, isPalindrome("aba"))
	assert.True(t, isPalindrome("A man, a plan, a canal: Panama"))
}
