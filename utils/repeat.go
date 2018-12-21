package utils

import "fmt"

// Repeatable is interface for Repeatable
type Repeatable interface {
	Do() error
}

// RepeatN repeat n times of Repeatable
func RepeatN(repeatable Repeatable, times uint) {
	for i := 0; i < int(times); i++ {
		err := repeatable.Do()
		if err != nil {
			fmt.Printf("%d - error - %+v\n", i, err)
		}
	}
}
