package main

import (
	"fmt"
	"time"
)

func emit(wordChannel chan string, done chan bool) {
	words := []string{"The", "quick", "brown", "fox", "jumped", "over", "the", "lazy", "dog"}
	i := 0

	t := time.NewTimer(3 * time.Second)

	// defer close(wordChannel)

	for {
		select {
		case wordChannel <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}
		case <-done:
			fmt.Println("All done! Terminating!")
			// close(done)
			done <- true
			return
		case <-t.C:
			// done <- true
			fmt.Println("Sending EOC to word channel.")
			wordChannel <- "EOC"
			// close(wordChannel)
			// close(done)
			return
		}
	}
}

func main() {
	wordCh := make(chan string)
	doneCh := make(chan bool)

	// This is for experimentation.
	//
	// One general principle of using Go channels is don't close a channel
	// from the receiver side. Because as you can see it complicates things dearly.
	// Also don't close a channel if the channel has multiple concurrent senders.
	//
	// In other words, you should only close a channel in a sender goroutine
	// And you should close a channel in the sender goroutine if and only if
	// the sender is the only sender or the last active sender of the channel.

	go emit(wordCh, doneCh)

	//for i := 0; i < 100; i++ {
	for word := range wordCh {
		if word == "EOC" {

			// since `emit()` has returned, no goroutine is able to receive from this channel.
			// wordCh <- "test"

			fmt.Println("Got EOC from word channel")
			close(wordCh)
			fmt.Println("Closed word channel")
			// doneCh <- true
			// break
		}
		fmt.Printf("%s\n", word)
	}

	fmt.Println("Sending true to done channel")
	// simply `doneCh <- true` will not work.
	// because `emit()` has exited, and there is no goroutine to receive
	// from the channel.
	go func() { doneCh <- true }()

	//go select {
	//case doneCh <- true:
	//	fmt.Println("Bello")
	//}

	// Program will not terminate until it receives something on this done channel.
	<-doneCh

	// close(wordCh)

	fmt.Println("Got stuff from the done channel")

	close(doneCh)

	fmt.Println("Closed done channel")
}
