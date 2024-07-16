// run without compilation with: go run hello.go
// run with compilation with: go build hello.go
// and then: ./hello
package main

import (
	"fmt"
)

// constants grouped together instead of declared separately
const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// private (not exported) functions start with a lowercase letter
// initializes variable prfix of type string, that is the return value
func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	// do not have to return prefix explicitly here
	return
}

// public (exported) functions start with a capital letter
func GetGreeting(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return getGreetingPrefix(language) + name
}

func main() {
	fmt.Println(GetGreeting("Tim", ""))
}
