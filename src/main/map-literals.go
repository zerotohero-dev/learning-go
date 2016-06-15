package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.0, -74.0,
	},
	"Google": Vertex{
		37.0, -122.0,
	},
}

var m2 = map[string]Vertex{
	"Bell Labs": {40.0, -74.0},
	"Google": {37.0, -122.0},
}

func main() {
	fmt.Println(m, m2)
}
