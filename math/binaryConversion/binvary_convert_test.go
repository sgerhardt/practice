package binaryConversion

import "testing"
import "github.com/stretchr/testify/assert"

func TestBinaryConvert(t *testing.T) {
	assert.Equal(t, `10101`, addBinary("1010", "1011"))
	assert.Equal(t, `100`, addBinary("11", "1"))
	assert.Equal(t, `11110`, addBinary("1111", "1111"))
}
