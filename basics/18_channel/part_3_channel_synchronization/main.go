package main

import (
	"fmt"
	"time"
)

/*
Use channels to synchronize execution accross goroutines
*/

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(2 * time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done // blocks until we get the notification from the worker function
}
