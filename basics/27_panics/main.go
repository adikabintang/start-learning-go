package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("panic testing")
	panic("a problem")

	_, err := os.Create("asu")
	if err != nil {
		fmt.Println("asuu")
		panic(err)
	}
}
