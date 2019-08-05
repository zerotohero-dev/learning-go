package main

import (
	"fmt"
	"time"
)

func printer(msg string, goCh chan bool) {
	<-goCh

	fmt.Printf("%s\n", msg)
}

func printerSelect(msg string, closeCh chan bool) {
	for {
		select {
		case <-closeCh:
			return
		default:
			fmt.Printf("%s\n", msg)
		}
	}
}

func main() {
	fmt.Println("Hello, world.")

	closeCh := make(chan bool)

	go printerSelect("bello", closeCh)

	time.Sleep(5 * time.Second)
	closeCh <- true
	time.Sleep(5 * time.Second)

	goCh := make(chan bool)

	// when something is sent to the channel, only one routine will receive it.
	// it is guaranteed to be delivered to only a single worker.
	// it is not a broadcast.
	// A broadcast mechanism does not exist within CSP channel mechanism itself;
	// you can fake one by using `close` since `close` will be dispatched to all receivers.
	for i := 0; i < 10; i++ {
		go printer(fmt.Sprintf("Printer %d", i), goCh)
	}

	time.Sleep(5 * time.Second)
	close(goCh)
	time.Sleep(5 * time.Second)
}
