package main

import (
	"fmt"
)

type Counter struct {
	value int
}

func (c *Counter) Inc() {
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func main() {
	fmt.Printf("main function in sync\n")
}
