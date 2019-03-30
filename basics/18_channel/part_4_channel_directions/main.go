package main

import (
	"fmt"
)

// variable_name chan<- type is only for receiving
func ping(pings chan<- string, msg string) {
	pings <- msg // receive
	// if we stry to send value to channel pings, it will cause compile error
}

// variable_name <- chan type is only for sending
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
