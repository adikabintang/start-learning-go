package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	//fmt.Println("received: ", <-c2) // wait for c2 first. if c1 is ready, it still blocks
	//fmt.Println("received: ", <-c1)

	for i := 0; i < 2; i++ {
		select { // wait for multiple channel communications. with select, it will block to wait for both c1 and c2 at the same time
		case msg2 := <-c2:
			fmt.Println("received: ", msg2)
		case msg1 := <-c1:
			fmt.Println("received: ", msg1)
		}
	}
}
