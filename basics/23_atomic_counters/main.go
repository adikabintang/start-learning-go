package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64 // counter

	// simulate 50 goroutines incrementing the counter about once a millisecond
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1) // atomically add the counter with 1

				time.Sleep(1 * time.Millisecond)
			}
		}()
	}

	time.Sleep(1 * time.Second) // wait for 1 second, see how counter variable ops will become

	opsFinal := atomic.LoadUint64(&ops) // safely get the ops counter variable by making a copy
	fmt.Println("ops: ", opsFinal)
}
