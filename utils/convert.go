package utils

import "strconv"

// Str2int convert string to int
func Str2int(str string) int {
	out, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return out
}

// Str2int64 convert string to int64
func Str2int64(str string) int64 {
	out, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return out
}

// Str2uint64 convert string to uint64
func Str2uint64(str string) uint64 {
	out, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return out
}
