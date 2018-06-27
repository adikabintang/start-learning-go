package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

func main() {
	a, e := f1(42)
	if e == nil {
		fmt.Println("a: ", a)
	} else {
		fmt.Println("Error: ", e)
	}
}
