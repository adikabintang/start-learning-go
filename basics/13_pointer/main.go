package main

import (
	"fmt"
)

func passByValue(val int) {
	val++
}

func passByPointer(val *int) {
	(*val)++
}

func main() {
	i := 1

	fmt.Println("i: ", i)
	passByValue(i)
	fmt.Println("i: ", i)
	passByPointer(&i)
	fmt.Println("i: ", i)
}
