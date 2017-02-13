package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Book struct {
	Title string `json:"title"`
	Year  int    `json:"published,omitempty"`
}

var data = []Book{
	{Title: "A Hitchhiker’s Guide to the Galaxy", Year: 2004},
}

func main() {
	jsn, err := json.Marshal(data)

	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}

	fmt.Println(data)
	fmt.Printf("%T %s\n", jsn, jsn)

	var titles []struct{ Title string }

	if err := json.Unmarshal(jsn, &titles); err != nil {
		log.Fatalf("JSON unmarshalling failed: %s", err)
	}

	fmt.Println(titles)
}
