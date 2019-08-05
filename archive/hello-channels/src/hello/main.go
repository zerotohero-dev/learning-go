package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working…")
	time.Sleep(time.Second)
	fmt.Println("done.")

	done <- true
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages

	fmt.Println(msg)

	// By default sends and received block until both
	// sender and receiver is ready.

	// Normally channels are unbuffered; i.e., they accept
	// sends ( chan <- ) only if there is a corresponding receive
	// ( <- chan ).
	//
	// Buffered channels can accept a limited number of values
	// without needing a corresponding receiver.

	messages2 := make(chan string, 2)

	messages2 <- "buffered"
	messages2 <- "channel"
	// messages2 <- "channel" // <- this will panic!

	fmt.Println(<-messages2)
	fmt.Println(<-messages2)

	done := make(chan bool, 1)
	go worker(done)

	// If you remove this line, the problem will exit without
	// even starting the worker.
	<-done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// You can still received remaining values inside a channel,
	// even if you closed it.
	for elem := range queue {
		fmt.Println(elem)
	}
}
