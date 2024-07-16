// run without compilation with: go run hello.go
// run with compilation with: go build hello.go
// and then: ./hello
package main

import (
	"fmt"
)

func getHelloWorld() string {
	return "Hello, world"
}

func main() {
	fmt.Println(getHelloWorld())
}
