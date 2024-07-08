package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddStrings(t *testing.T) {
	assert.Equal(t, "4", addStrings("2", "2"))
	assert.Equal(t, "12", addStrings("5", "7"))
}
