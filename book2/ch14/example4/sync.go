package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	// mutex mutual exclusion lock, initialized with unlocked mutex (zero value)
	// here mutex embeded into the struct
	// it works, but the book suggests to avoid it and give a named field, like
	// mu: sync.Mutex
	// from example3
	sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func main() {
	fmt.Printf("main function in sync\n")
}
