package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]

	if edges == nil {
		edges = make(map[string]bool)

		graph[from] = edges
	}

	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("lorem", "dolor")
	addEdge("dolor", "sit")
	addEdge("sit", "lorem")

	fmt.Println(hasEdge("sit", "lorem"))
}
