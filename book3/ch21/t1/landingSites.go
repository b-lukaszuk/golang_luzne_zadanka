// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	Lat  float64
	Long float64
}

// prints any erros and exits
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It displays the JSON endocing of the three rover landing sites.\n\n")

	curiosity := location{-4.5895, 137.4417}
	spirit := location{-14.5684, 175.472636}
	opportunity := location{-1.9462, 354.4734}

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)
	fmt.Printf("curiosity: %s\n", string(bytes))

	bytes, err = json.Marshal(spirit)
	exitOnError(err)
	fmt.Printf("spirit: %s\n", string(bytes))

	bytes, err = json.Marshal(opportunity)
	exitOnError(err)
	fmt.Printf("opportunity: %s\n", string(bytes))

	fmt.Printf("\nThat's all. Goodbye!\n")
}
