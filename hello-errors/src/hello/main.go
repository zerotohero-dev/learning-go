package main

import (
	"errors"
	"fmt"
	"os"
)

// In go, it is idiomatic to communicate errors with a separate
// explicit return value. That value is, by convention, the second
// return value from the function. Also, by conventions, errors
// have type `error`, which is a builtin interface.

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can’t work with 42")
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d — %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	a1 := argError{1, "foo"}
	a2 := argError{2, "bar"}

	fmt.Println(a1.Error())
	fmt.Println((&a2).Error())

	if arg == 42 {
		return -1, &argError{arg, "can’t work with it"}
		// This won’t work.
		// 		return -1, argError{arg, "can’t work with it"}
		// So, go automatically dereferences the struct when you
		// call a method on a struct; however, when you return
		// a struct to implement an interface you need to
		// dereference explicitly.
	}

	return arg + 3, nil
}

func main() {
	if n, err := fmt.Printf("Hello, World!\n"); err != nil {
		fmt.Println(n)
		os.Exit(1)
	} else {
		fmt.Printf("Printed %d bytes\n", n)
	}

	// fmt.Printf("Printed %d bytes\n", n)

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
