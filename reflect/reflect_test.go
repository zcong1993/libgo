package reflecthelper

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type f1 struct {
	id string
}

type f2 struct {
	f1 *f1
}

func TestIndirectType(t *testing.T) {
	type fixtures []struct {
		in   reflect.Type
		want reflect.Kind
	}

	fs := fixtures{
		{in: reflect.TypeOf(&f1{}), want: reflect.Struct},
		{in: reflect.TypeOf([]f2{}), want: reflect.Struct},
	}

	for _, f := range fs {
		assert.Equal(t, IndirectType(f.in).Kind(), f.want)
	}
}

func TestSlice2Map(t *testing.T) {
	type user struct {
		ID string
	}

	a := user{ID: "a"}
	b := user{ID: "b"}

	users := []user{a, b}

	out := map[interface{}]interface{}{
		"a": a,
		"b": b,
	}

	assert.Equal(t, Slice2Map(users, "ID"), out)
}
