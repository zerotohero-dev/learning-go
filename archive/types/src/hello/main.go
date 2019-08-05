package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// You can attach functions, that are called “methods” to a type;
// which is, kind of, closer to object-orientation.

type webPage struct {
	url  string
	body []byte
	err  error
}

type summableSlice []int

func (s summableSlice) sum() int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}

func (w *webPage) get() {
	// fmt.Println("Getting…")
	// fmt.Printf("%v\n", w)

	res, err := http.Get(w.url)

	if err != nil {
		// fmt.Println("err001")
		w.err = err

		return
	}
	// else {
	//	fmt.Println("No err 001")
	// }

	defer res.Body.Close()

	w.body, err = ioutil.ReadAll(res.Body)

	fmt.Println(len(w.body))

	if err != nil {
		// fmt.Println("err002")
		w.err = err
	}
	// else {
	//	fmt.Println("No err 002")
	// }
}

func (w *webPage) isOK() bool {
	// fmt.Println("called isOK")
	// fmt.Println(w.err)
	// fmt.Println("+++++++++++++")
	return w.err == nil
}

func main() {
	// 		w := webPage{url: "https://bytesized.tv"}
	// This will also work just fine because go will pass a pointer to
	// the struct abouve, into the receiver realizing we actually want to
	// use a pointer type.
	// But to be explicit, it is better to use the address of the object.
	// 		w.get() // equivalent to `{&w}.get()
	//
	// So this usage is better. — Being explicit is generally better.
	w := &webPage{url: "https://bytesized.tv"}

	w.get()

	if w.isOK() {
		fmt.Println("Everything is awesome!")
		fmt.Printf("body length: %d", len(w.body))
		fmt.Printf("results: %s %v %d", w.url, w.err, len(w.body))
	} else {
		fmt.Println("Things are not okay!")
		fmt.Println(w.isOK())
	}

	fmt.Println("Something went wrong")
	fmt.Println(w.err)

	s := summableSlice{1, 2, 3, 4, 5, 5, 5, 4, 1}

	fmt.Println("sum", s.sum())
}
