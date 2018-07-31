package main

import (
	"fmt"
	"time"
)

/*
Rate limiting is a mechanism for controlling resource utilization and maintaining quality of service
*/

func main() {

	//////////////// BASIC RATE LIMITING
	requests := make(chan int, 5)

	// send 5 requests without delay
	for i := 1; i < 5; i++ {
		requests <- i
	}
	close(requests)

	// limit to receive per 200 ms
	limiter := time.Tick(200 * time.Millisecond) // Tick is a convenience wrapper for NewTicker providing access to the ticking channel only

	for req := range requests {
		<-limiter
		fmt.Println("request ", req, time.Now())
	}

	////////////////// END OF BASIC RATE LIMITING

	////////////////// RATE LIMT BURST REQUEST
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		// send to burstyLimiter every 200 ms
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// simulate 5 more requests
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request ", req, time.Now())
	}
}
