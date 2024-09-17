// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It displays error returned by url.Parse.\n\n")

	correctUrl := "https://google.pl"
	response, err := url.Parse(correctUrl)
	if err == nil {
		fmt.Printf("Testing '%s', response short:\n%v\n", correctUrl, response)
		fmt.Printf("Testing '%s', response long\n%#v\n", correctUrl, response)
	} else {
		fmt.Printf("Testing '%s', error short:\n%v\n", correctUrl, response)
		fmt.Printf("Testing '%s', error long:\n%v\n", correctUrl, response)
		os.Exit(1)
	}

	wrongUrl := "https://w h a t e v e r.com"
	response, err = url.Parse(wrongUrl)
	if err == nil {
		fmt.Printf("Testing '%s', response short:\n%v\n", wrongUrl, response)
		fmt.Printf("Testing '%s', response long\n%#v\n", wrongUrl, response)
	} else {
		fmt.Printf("Testing '%s', error short:\n%v\n", wrongUrl, response)
		fmt.Printf("Testing '%s', error long:\n%v\n", wrongUrl, response)

		if e, ok := err.(*url.Error); ok {
			fmt.Printf("Testing '%s', error very long:\n%v\n", wrongUrl, e.Err)
		}

		os.Exit(1)
	}

	fmt.Printf("\nThat's all. Goodbye!\n")
}
