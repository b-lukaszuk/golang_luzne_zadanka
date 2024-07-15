package invoice

import (
	"fmt"
)

var a = func() string {
	fmt.Println("variable a initialized")
	return "value of a"
}()

var b = func() string {
	fmt.Println("variable b initialized", a)
	return "value of b"
}()

var c = func() string {
	fmt.Println("variable c initialized", b)
	return "value of c"
}()

func init() {
	fmt.Println("Running invoice init", c)
}

func Print() {
	fmt.Println("INVOICE Number X")
}
