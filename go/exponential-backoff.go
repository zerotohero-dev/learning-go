package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute

	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)

		if err == nil {
			return nil
		}

		log.Printf("server not responding (%s); retrying…", err)

		time.Sleep(time.Second << uint(tries))
	}

	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	if err := WaitForServer("https://volkan.io/"); err != nil {

		// Same as `log.Fatalf`
		fmt.Fprintf(os.Stderr, "Site is down: %v", err)
		os.Exit(1)
	}
}
