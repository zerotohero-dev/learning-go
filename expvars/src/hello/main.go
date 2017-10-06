package main

import (
	"expvar"
	"net/http"
)

func main() {
	views := expvar.NewInt("views")

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		views.Add(1)
	})

	http.ListenAndServe(":8888", nil)
}
