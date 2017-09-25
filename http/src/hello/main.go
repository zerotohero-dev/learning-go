package main

import (
	"fmt"
	"net/http"
)

func poemHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	names := r.Form["name"]

	if len(names) == 0 {
		http.Error(w, "Poop!", http.StatusBadRequest)

		return
	}

	poemName := names[0]

	fmt.Fprintf(w, "Poem %s coming soon\n", poemName)
}

func main() {
	http.HandleFunc("/poem", poemHandler)
	http.ListenAndServe(":8080", nil)
}
