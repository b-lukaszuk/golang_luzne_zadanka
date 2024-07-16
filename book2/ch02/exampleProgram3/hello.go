// run without compilation with: go run hello.go
// run with compilation with: go build hello.go
// and then: ./hello
package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "

func getGreeting(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(getGreeting("Tim"))
}
