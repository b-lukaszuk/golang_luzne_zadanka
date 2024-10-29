// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
)

func add2Channel(words *[]string, c chan string) {
	for _, w := range *words {
		c <- w
	}
	close(c)
}

func filterOutConseqDuplicates(upstream, downstream chan string) {
	prevItem := ""
	for item := range upstream {
		if item != prevItem {
			downstream <- item
		}
		prevItem = item
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for item := range upstream {
		fmt.Printf("%s\n", item)
	}
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It uses pipelines to remove consecutive duplicates from list of words.\n")

	words := []string{"a", "b", "b", "b", "c", "d", "c", "c", "z", "y", "y"}
	c0 := make(chan string)
	c1 := make(chan string)

	fmt.Printf("List of words with consecutive duplicates: %s\n", words)
	fmt.Printf("List after consecutive duplicates have been removed:\n")
	go add2Channel(&words, c0)
	go filterOutConseqDuplicates(c0, c1)
	printGopher(c1)

	fmt.Printf("\nThat's all. Goodbye!\n")
}
