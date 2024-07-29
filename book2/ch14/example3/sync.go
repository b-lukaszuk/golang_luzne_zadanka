package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	// mutex mutual exclusion lock, initialized with unlocked mutex (zero value)
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func main() {
	fmt.Printf("main function in sync\n")
}
