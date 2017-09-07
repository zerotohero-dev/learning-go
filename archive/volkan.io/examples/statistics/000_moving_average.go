package main

import "fmt"

type MovingAverage struct {
	window []int
	length int
	insert int
	sum int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		window: make([]int, size),
		length: 0,
		insert: 0,
		sum: 0,
	}
}

func (avg *MovingAverage) Next(val int) float64 {
	if avg.length < len(avg.window) {
		avg.length++
	}

	avg.sum -= avg.window[avg.insert]
	avg.sum += val
	avg.window[avg.insert] = val

	avg.insert = (avg.insert + 1) % len(avg.window)

	return float64(avg.sum)/float64(avg.length)
}

func main() {
	obj := Constructor(2)

	fmt.Println( obj.Next(10) )
	fmt.Println( obj.Next(13) )
	fmt.Println( obj.Next(21) )
	fmt.Println( obj.Next(44) )
}