package main

import (
	"fmt"
	"log"
)

type FooError struct {
}

func (f *FooError) Error() string {
	return "Boom goes the dynomite!"
}

var ErrBar = errors.New("Bar!!")

func Boom() error {
	return fmt.Errorf("Boom goes %s", "the dynomite!")
}

func Boom2() error {
	return &FooError{}
}

func Boom3() error {
	return ErrBar
}

func main() {
	err := Boom()

	if err != nil {
		log.Println(err)
	}

	err = Boom2()

	if err != nill {
		log.Println(err)
	}

	err = Boom3()

	if err == ErrBar {
		log.Println("Got the bar!")
	}

	switch e := err.(type) {
	case *FooError:
		log.Println("Foo")
	default:
		if err == ErrBar {
			log.Println("Bar")
		}
	}
}
