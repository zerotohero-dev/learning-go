package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 3200)

	for req := range requests {
		<-limiter
		fmt.Println("request1", req, time.Now())
	}

	burstLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 3200) {
			burstLimiter <- t
		}
	}()

	burstRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstRequests <- i
	}
	close(burstRequests)

	for req := range burstRequests {
		<-burstLimiter
		fmt.Println("request2", req, time.Now())
	}
}
