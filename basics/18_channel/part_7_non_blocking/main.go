package main

import (
	"fmt"
)

func main() {
	messagesChan := make(chan string)
	signalsChan := make(chan bool)

	// non-blocking receive. If there is no one sends through channel, do what in the "default"
	select {
	case msgStr := <-messagesChan:
		fmt.Println("received: ", msgStr)
	default:
		fmt.Println("no message received")
	}

	// non-blocking send. channel has no buffer and no receiver, so it cannot send to the channel
	msg := "hi"
	select {
	case messagesChan <- msg:
		fmt.Println("sent message: ", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messagesChan:
		fmt.Println("received message", msg)
	case sig := <-signalsChan:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
