package main

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s\n", name)
}

func main() {
	fmt.Print("test")
}
