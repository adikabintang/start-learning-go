package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 2) // a channel of strings buffering up to 2 values

	messages <- "buffered" // send value to channel without corresponding receive (without var <- messages)
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
