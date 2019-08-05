package main

import "fmt"

func emit(c chan string) {
	words := []string{"quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}

	for _, word := range words {
		c <- word
	}

	// close(c)
}

func main() {
	wordChannel := make(chan string)

	go emit(wordChannel)
	go emit(wordChannel)

	someWord := <- wordChannel

	fmt.Printf("%s\n", someWord)

	// receive stuff from channel until the channel is closed:
	for word := range wordChannel {
		fmt.Printf("%s\n", word)
	}

	someWord = <- wordChannel

	fmt.Printf("%s\n", someWord)


	someWord, ok := <- wordChannel

	// receive from Channel: <-ch
	// send to Channel: ch<- x

	if ok != true {
		fmt.Println("Channel is closed buddy!")
	}

	fmt.Println()
}
