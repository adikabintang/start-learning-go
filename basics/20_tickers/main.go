package main

import (
	"fmt"
	"time"
)

/*
Timer: trigger only once
Ticker: trigger at that interval
*/

func main() {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at: ", t)
		}
	}()

	time.Sleep(3000 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
