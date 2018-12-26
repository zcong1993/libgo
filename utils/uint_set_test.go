package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUintSet_Add(t *testing.T) {
	us := NewUintSet()
	us.Add(1)
	us.Add(2)
	us.Add(3)
	us.Add(1)

	assert.Equal(t, len(us), 3)
}

func TestUintSet_Adds(t *testing.T) {
	us := NewUintSet()
	us.Adds(1, 2, 3, 1)

	assert.Equal(t, len(us), 3)
}

func TestUintSet_ToSlice(t *testing.T) {
	us := NewUintSet(1)
	us.Add(1)
	us.Add(2)
	us.Add(3)
	us.Add(1)

	assert.Equal(t, len(us), 3)
}
