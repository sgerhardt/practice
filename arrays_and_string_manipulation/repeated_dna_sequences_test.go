package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindRepeated(t *testing.T) {
	assert.Equal(t, []string{"AAAAACCCCC", "CCCCCAAAAA"}, findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	assert.Equal(t, []string{"AAAAAAAAAA"}, findRepeatedDnaSequences("AAAAAAAAAAA"))
}
