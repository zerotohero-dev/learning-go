package main

import (
	"time"
	"fmt"
)

func emit(channelCh chan chan string, done chan bool) {
	wordChannel := make(chan string)
	channelCh <- wordChannel

	defer close(wordChannel)
	defer close(done)
	defer close(channelCh)

	words := []string{"The", "quick", "brown", "fox", "jumped", "over", "the", "lazy", "dog"}

	i := 0

	t := time.NewTimer(3 * time.Second)

	for {
		select {
		case wordChannel <- words[i]:
			i++
			if i == len(words) {
				i = 0
			}

		case <-done:
			done <- true

			return
		case <-t.C:
			return
		}
	}
}

func main() {
	// wordCh := make(chan string)
	channelCh := make(chan chan string)
	doneCh := make(chan bool)

	go emit(channelCh, doneCh)

	wordCh := <-channelCh

	for word := range wordCh {
		fmt.Printf("%s ", word)
	}
}
