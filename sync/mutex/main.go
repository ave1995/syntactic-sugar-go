package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	c := SafeCounter{}

	for i := 0; i < 1000; i++ {
		c.Inc()
	}

	time.Sleep(time.Second)

	fmt.Println("count:", c.Value())
}
