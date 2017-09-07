package safe

import (
    "sync"
)

type Counter struct {
    v map[string]int
    mux sync.Mutex
}

func (c *Counter) Inc(key string) {
    c.mux.Lock()
    c.v[key]++
    c.mux.Unlock()
}

func (c *Counter) Value(key string) int {
    c.mux.Lock()
    defer c.mux.Unlock()
    return c.v[key]
}

func NewCounter() Counter {
    return Counter{v: make(map[string]int)}
}