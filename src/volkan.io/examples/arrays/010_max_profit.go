package main

import "fmt"

func maxProfit(prices []int) int {
	total := 0

	for i := 0; i < len(prices) - 1; i++ {
		if prices[i+1] > prices[i] {
			total += prices[i+1] - prices[i]
		}
	}

	return total
}

func main() {
	fmt.Println(
		maxProfit([]int{10,22,33,12,22,11,9,8,56,11,4,92}),
	)
}

