package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "message"
	}()

	// "select" will wait for both c1 and time.After and do whatever comes first
	select {
	case res := <-c1:
		fmt.Println("c1: ", res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
