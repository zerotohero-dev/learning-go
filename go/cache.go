package main

import (
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()

	return v
}

func main() {
	cache.mapping["lorem"] = "ipsum"

	Lookup("lorem")
}
