package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type config struct {
	Route       string
	BindAddress string `json:"addr"`
}

func main() {
	f, err := os.Open("config.json")

	if err != nil {
		os.Exit(1)
	}

	var c config

	dec := json.NewDecoder(f)

	err = dec.Decode(&c)

	f.Close()

	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	fmt.Println("Hello ", c.Route)

	http.HandleFunc(c.Route, func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)

		meaning := map[string]int{
			"life":     42,
			"universe": 0,
		}

		enc.Encode(meaning)
	})
	http.ListenAndServe(c.BindAddress, nil)
}
