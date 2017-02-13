package main

// go get golang.org/x/net/html
// go build functions.go
// go build fetch.go
// ./fetch https://volkan.io/ | ./functions

import (
	"fmt"
	"golang.org/x/net/html"
	"math"
	"os"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func main() {
	doc, err := html.Parse(os.Stdin)

	hypot(3, 4)

	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)

		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}
