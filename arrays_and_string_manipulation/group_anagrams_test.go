package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	assert.Equal(t, [][]string{{"eat", "tea"}, {"tan"}}, groupAnagrams([]string{"eat", "tea", "tan"}))
	assert.Equal(t, [][]string{{"bdddddddddd"}, {"bbbbbbbbbbc"}}, groupAnagrams([]string{"bdddddddddd", "bbbbbbbbbbc"}))
}
