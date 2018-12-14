package utils

import "fmt"

type Repeatable interface {
	Do() error
}

func RepeatN(repeatable Repeatable, times uint) {
	for i := 0; i < int(times); i++ {
		err := repeatable.Do()
		if err != nil {
			fmt.Printf("%d - error - %+v\n", i, err)
		}
	}
}
