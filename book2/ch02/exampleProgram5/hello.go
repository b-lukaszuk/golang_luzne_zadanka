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
	if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(getGreeting("Tim", ""))
}
