package main

import "fmt"




func printer(msg, msg2 string, repeat int, msgs ...string) (string, error) {
	defer fmt.Printf("Geronimo 1!\n")
	defer fmt.Printf("Geronimo 2!\n")

	for _, msg := range msgs {
		fmt.Printf("%s\n", msg)
	}

	for repeat > 0 {
		_, err := fmt.Printf("%s %s\n", msg, msg2)

		if err != nil {
			return msg, err
		}

		repeat--
	}

	return msg, nil
}

func main() {
	// fmt.Printf("%q", foo) => print in a “friendly” manner.
	// %x: hex version.
	printer("Hello", ", World!", 3, "lumos maxima!")
}
