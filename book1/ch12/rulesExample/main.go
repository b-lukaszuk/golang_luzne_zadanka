package main

import (
	"fmt"
	"rulesExample/invoice"
)

func init() {
	fmt.Println("Running init for main")
}

func main() {
	fmt.Println("--program start--")
	invoice.Print()
}
