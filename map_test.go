package jgunify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeysOfMap(t *testing.T) {
	myMap := map[string]int{"one": 1, "two": 2, "three": 3}
	keys := KeysOfMap(myMap)
	assert.Equal(t, []string{"one", "two", "three"}, keys)
}
