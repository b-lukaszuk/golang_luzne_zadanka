// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
)

type item struct {
	name string
}

type character struct {
	name     string
	leftHand *item
}

func (c *character) String() string {
	var article string = "\b"
	if c.leftHand != nil {
		article = "a/an"
	}
	return fmt.Sprintf("%s with %s %s in his left hand", c.name, article, c.leftHand)
}

func (i *item) String() string {
	if i == nil {
		return "nothing"
	}
	return i.name
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("%s picks up a/an %s.\n", c, i)
	c.leftHand = i
}

func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}
	if c.leftHand == nil {
		fmt.Printf("%s has %s to give to %s.\n", c, c.leftHand, to)
	}
	if to.leftHand != nil {
		fmt.Printf("%s cannot take a/an %s, he got his left hand full.\n", to, c.leftHand)
	}
	fmt.Printf("%s gives the %s to %s.\n", c, c.leftHand, to)
	to.leftHand = c.leftHand
	c.leftHand = nil
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It defines characters and items and displays their short history.\n\n")

	shield := &item{"shield"}
	arthur := &character{"Arthur", nil}
	knight := &character{"Unknown Knight", nil}

	arthur.pickup(shield)
	arthur.give(knight)
	fmt.Printf("Now, there is %s.\n", knight)

	fmt.Printf("\nThat's all. Goodbye!\n")
}
