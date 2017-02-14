package main

import (
    "net/http"
    "log"
    "fmt"
)

func main() {
    db := database{"shoes": 50, "socks": 5}

    mux := http.NewServeMux()

    mux.Handle("/list", http.HandlerFunc(db.list))
    //mux.Handle("/price", http.HandlerFunc(db.price))
    mux.HandleFunc("/price", db.price)

    log.Fatal(http.ListenAndServe("0.0.0.0:8080", mux))

    // Using DefaultServeMux
    // db := database{" shoes": 50, "socks": 5}
    // http.HandleFunc("/ list", db.list)
    // http.HandleFunc("/ price", db.price)
    // log.Fatal( http.ListenAndServe(" localhost: 8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
        for item, price := range db {
            fmt.Fprintf(w, "%s: %s\n", item, price)
        }
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
        item := req.URL.Query().Get("item")
        price, ok := db[item]

        if !ok {
            w.WriteHeader(http.StatusNotFound)
            fmt.Fprintf(w, "no such item: %q", item)

            return
        }

        fmt.Fprintf(w, "%s\n", price)
}