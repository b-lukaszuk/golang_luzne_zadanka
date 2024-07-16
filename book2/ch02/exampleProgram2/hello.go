// run without compilation with: go run hello.go
// run with compilation with: go build hello.go
// and then: ./hello
package main

import (
	"fmt"
)

func getGreeting(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(getGreeting("Tim"))
}
