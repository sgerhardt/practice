package hashing_and_hash_tables

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDesignHashMap(t *testing.T) {
	myMap := Constructor()
	myMap.Put(1, 5)
	assert.Equal(t, 5, myMap.Get(1))
	myMap.Remove(1)
	assert.Equal(t, -1, myMap.Get(1))
}
