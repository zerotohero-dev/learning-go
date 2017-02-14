package main

import (
    "net/http"
    "log"
    "fmt"
)

type dollars float32
type database map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    switch req.URL.Path {
    case "/list":
        for item, price := range db {
            fmt.Fprintf(w, "%s: %s\n", item, price)
        }
    case "/price":
        item := req.URL.Query().Get("item")
        price, ok := db[item]

        if !ok {
            w.WriteHeader(http.StatusNotFound)
            fmt.Fprintf(w, "no such item: %q", item)

            return
        }

        fmt.Fprintf(w, "%s\n", price)
    default:
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "no such page: %q", req.URL)
    }

}

func main() {
    db := database{"shoes": 44, "socks": 55}

    log.Fatal(http.ListenAndServe("0.0.0.0:8080", db))
}