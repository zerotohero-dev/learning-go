package main

import "fmt"

func canWinNim(n int) bool {
	return n % 4 != 0
}

func main() {
	fmt.Println(canWinNim(43))
}
