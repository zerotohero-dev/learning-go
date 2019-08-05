package main

import "fmt"

type MemoizeFn func(int, ...int) interface{}

var Fibonacci MemoizeFn

func init() {
	Fibonacci = Memoize(func(x int, xs ...int) interface{} {
		if x < 2 {
			return x
		}

		return Fibonacci(x-1).(int) + Fibonacci(x-2).(int)
	})
}

func Memoize(fn MemoizeFn) MemoizeFn {
	cache := make(map[string]interface{})

	return func(x int, xs ...int) interface{} {
		key := fmt.Sprint(x)

		for _, i := range xs {
			key += fmt.Sprintf(",%d", i)
		}

		if value, found := cache[key]; found {
			return value
		}

		value := fn(x, xs...)
		cache[key] = value

		return value
	}
}

func main() {
	fmt.Println(Fibonacci(6))
	fmt.Println(Fibonacci(6))
}
