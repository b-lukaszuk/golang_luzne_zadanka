// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
	"strings"
)

func add2Channel(words []string, c chan string) {
	for _, w := range words {
		c <- w
	}
	close(c)
}

func printGopher(upstream chan string) {
	for item := range upstream {
		fmt.Printf("%s\n", item)
	}
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It uses pipelines to split string to words and print them on screen.\n")

	txt := "life is like a hurricane here in Duckburg"
	c0 := make(chan string)

	fmt.Printf("Original text: '%s'\n", txt)
	fmt.Printf("The text split into words:\n")
	go add2Channel(strings.Fields(txt), c0)
	printGopher(c0)

	fmt.Printf("\nThat's all. Goodbye!\n")
}
