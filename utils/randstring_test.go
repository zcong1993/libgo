package utils_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/libgo/utils"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	a := make([]int, 10)
	l := 32
	for range a {
		r1, err := utils.GenerateRandomString(l)
		assert.Nil(t, err)
		assert.Equal(t, len(r1), l)
		r2, err := utils.GenerateRandomString(l)
		assert.Nil(t, err)
		assert.Equal(t, len(r2), l)

		assert.NotEqual(t, r1, r2)
	}
}
