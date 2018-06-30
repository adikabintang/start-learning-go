package main

import (
	"fmt"
)

// goroutine is a lightweight thread of execution

func f(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str, ": ", i)
	}
}

func main() {
	// a standard way to call a function
	f("blocking")

	// use go routing. It's like use a separate thread from main thread
	go f("with goroutines")

	// goroutine can also be used for anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("matt schofield")

	// hold the program until a button is pressed. like getchar() in C
	fmt.Scanln()
	fmt.Println("done")

}
