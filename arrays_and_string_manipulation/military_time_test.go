package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimeConversion(t *testing.T) {
	assert.Equal(t, "19:05:45", timeConversion("07:05:45PM"))
	assert.Equal(t, "00:40:22", timeConversion("12:40:22AM"))
}
