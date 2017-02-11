package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Println("error!")
				continue
			}

			countLines(f, counts)

			f.Close()
		}
	}

	// Note: Ignoring potential errors from input.Err()

	fmt.Println("----------")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// Although there is no set idiomatic way of doing this, it looks like public/important
// functions go to the top of the go file, and private/less-important functions go to the bottom.
// The ordering of functions in a module does not impact anything.

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		fmt.Println("Scan…")
		counts[input.Text()]++
		fmt.Println(counts, len(counts))

		if len(counts) > 3 {
			break
		}
	}
}
