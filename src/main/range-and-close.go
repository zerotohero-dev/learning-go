package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		time.Sleep(1000*time.Millisecond)
		c <- x
		x, y = y, x+y
	}

	// Only the sender should close a channel, never the receiver.
	// Sending on a closed channel will cause a panic.
	close(c)
}

func main() {
	c := make(chan int, 10)

	go fibonacci(cap(c), c)

	fmt.Println("Created the channel")

	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("After for…")
}
