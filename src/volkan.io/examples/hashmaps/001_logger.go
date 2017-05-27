package main

import (
	"fmt"
)

type Logger struct {
	Ok map[string]int
}

func Constructor() Logger {
	ok := make(map[string]int)



	return Logger{Ok: ok}
}

func (logger *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if timestamp < logger.Ok[message] {
		return false
	}

	logger.Ok[message] = timestamp + 10

	return true
}

func main() {
	obj := Constructor()
	param_1 := (&obj).ShouldPrintMessage(1 ,"Hello World")
	fmt.Println(param_1)
	param_1 = obj.ShouldPrintMessage(2 ,"Hello World")
	fmt.Println(param_1)
	param_1 = obj.ShouldPrintMessage(21 ,"Hello World")
	fmt.Println(param_1)
}