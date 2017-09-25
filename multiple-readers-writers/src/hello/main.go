package main

import (
	"net/http"
	"io/ioutil"
	"os"
	"fmt"
)

func getPage(url string) (int, error) {
	resp, err := http.Get(url)

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return 0, err
	}

	return len(body), nil
}

//func getter(url string, size chan string) {
//	length, err := getPage(url)
//
//	if err == nil {
//		size <- fmt.Sprintf("%s has length %d", url, length)
//	}
//}

func worker(urlCh chan string, sizeCh chan string, id int) {
	for {
		url := <-urlCh
		length, err := getPage(url)

		if err == nil {
			sizeCh <- fmt.Sprintf("%s has length %d (%d)", url, length, id)
		} else {
			sizeCh <- fmt.Sprintf( "Error getting %s: %s", url, err)
		}
	}
}

func generator(url string, urlCh chan string) {
	urlCh <- url
}

func main() {
	urls := []string{"https://bytesized.tv", "https://volkan.io", "https://o2js.com"}

	urlCh := make(chan string)
	sizeCh := make(chan string)

	for i := 0; i < 10; i++ {
		go worker(urlCh, sizeCh, i)
	}

	for _, url := range urls {
		// urlCh <- url
		go generator(url, urlCh)
	}

	//	urlCh <- urls[0]

	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n", <-sizeCh)
	}

	//size := make(chan string)
	//
	//for _, url := range urls {
	//	go getter(url, size)
	//	// length, err := getPage(url)
	//
	//	// if err != nil {
	//	// os.Exit(1)
	//	// }
	//	// fmt.Printf("%s is length %d\n", url, length)
	//}
	//
	//for i := 0; i < len(urls); i++ {
	//	fmt.Printf("%s\n", <-size)
	//}

	os.Exit(0)
}
