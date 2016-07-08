package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1

	for {
		time.Sleep(3000*time.Millisecond)
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}


func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("c:")
			fmt.Println(<-c)
			fmt.Println("________")
		}
		quit <- 0
	}()

	time.Sleep(3000*time.Millisecond)
	fmt.Println("fibo")
	time.Sleep(3000*time.Millisecond)
	fibonacci(c, quit)
	time.Sleep(3000*time.Millisecond)
}
