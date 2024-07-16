// run without compilation with: go run hello.go
// run with compilation with: go build hello.go
// and then: ./hello
package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const french = "French"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func getGreeting(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(getGreeting("Tim", ""))
}
