package dynamic_programming_and_recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPermutations(t *testing.T) {

	//expected := map[string]struct{}{"ab": {}, "ba": {}}
	//assert.Equal(t, expected, getPermutations("ab"))
	expected := map[string]struct{}{"act": {}, "atc": {}, "cat": {}, "cta": {}, "tac": {}, "tca": {}}
	assert.Equal(t, expected, getPermutations("cat"))
}
