// simple solution,
// not entirely idiot/error-proof, not optimized for speed
package main

import (
	"errors"
	"fmt"
)

func convertStringToBool(in string) (bool, error) {
	switch {
	case in == "true" || in == "yes" || in == "1":
		return true, nil
	case in == "false" || in == "no" || in == "0":
		return false, nil
	default:
		return false, errors.New("input not recognised")
	}
}

func declareConversion(in string) {
	var res, err = convertStringToBool(in)
	if err == nil {
		fmt.Printf("Converting %q to bool, result: %v\n", in, res)
	} else {
		fmt.Printf("Cannot convert %q to bool, conversion undefined\n", in)
	}
}

func main() {

	fmt.Printf("Toy program.\n")
	fmt.Printf("It converts strings to bools (if possible).\n\n")

	var tests []string = []string{"true", "false", "yes", "no", "1", "0", "a", "5"}
	for _, t := range tests {
		declareConversion(t)
	}

	fmt.Printf("\nThat's all. Goodbye!\n")
}
