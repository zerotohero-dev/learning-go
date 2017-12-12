package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("Creating…")

	f, err := os.Create(p)

	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing…")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("Closing…")
	f.Close()
}

func main() {
	f := createFile("/tmp/defer.txt")

	defer closeFile(f)
	writeFile(f)
}
