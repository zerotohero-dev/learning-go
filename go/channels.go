package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {

		// Start a goroutine:
		go fetch(url, ch)
	}

	for range os.Args[1:] {

		// Receive form channel:
		fmt.Println(<-ch)
	}

	fmt.Printf("%2fs elapsed", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)

		return
	}

	bytes, err := io.Copy(ioutil.Discard, resp.Body)

	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprint(err)

		return
	}

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs %7d %s", secs, bytes, url)
}
