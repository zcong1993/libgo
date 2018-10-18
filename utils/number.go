package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Num2String transfer number to string, will panic when given other type
func Num2String(in interface{}) string {
	switch reflect.TypeOf(in).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", in)
	case reflect.Float32, reflect.Float64:
		return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%f", in), "0"), ".")
	default:
		panic(errors.New("invalid type"))
	}
}
