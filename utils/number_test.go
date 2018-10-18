package utils_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/libgo/utils"
	"testing"
)

func TestNum2String(t *testing.T) {
	tt := assert.New(t)

	type fixtures []struct {
		want string
		in   interface{}
	}

	f := fixtures{
		{"10", int(10)},
		{"10", int8(10)},
		{"10", int16(10)},
		{"10", int32(10)},
		{"10", int64(10)},
		{"10", uint8(10)},
		{"10", uint16(10)},
		{"10", uint32(10)},
		{"10", float32(10.000)},
		{"10", float64(10.00)},
	}

	for _, v := range f {
		tt.Equal(v.want, utils.Num2String(v.in))
	}

	tt.Panics(func() {
		utils.Num2String("not a number")
	})
}
