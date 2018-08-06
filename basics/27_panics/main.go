package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("panic testing")

	/*
		When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller.
	*/
	panic("a problem")

	_, err := os.Create("asu")
	if err != nil {
		fmt.Println("asuu")
		panic(err)
	}
}
