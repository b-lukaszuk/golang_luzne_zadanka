// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
	"strings"
)

// text should not contain too many strange characters
// only removes commas and dots from text
func getCounts(text string) map[string]int {
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, ".", "")
	var words []string = strings.Fields(strings.ToLower(text))
	var counts map[string]int = make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	return counts
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It counts words in a text (case insensitive).\n\n")

	var text string = "a b A c d. car d, d Car"
	fmt.Printf("Original text: %q\n", text)
	fmt.Printf("Counts: %v\n", getCounts(text))

	fmt.Printf("\nThat's all. Goodbye!\n")
}
