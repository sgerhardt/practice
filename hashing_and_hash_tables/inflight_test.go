package hashing_and_hash_tables

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFindMoviesForFlightDuration(t *testing.T) {
	flightLength := 200
	movieLengths := []int{50, 200, 150}
	assert.True(t, canFindMoviesForFlightDuration(flightLength, movieLengths))
}
