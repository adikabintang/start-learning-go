package main

import (
	"fmt"
)

/*
Channels are the pipe connecting goroutines.
we can send a value from goroutine to another
*/

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	// it blocks until messages arrives
	msg := <-messages
	fmt.Println(msg)
}
