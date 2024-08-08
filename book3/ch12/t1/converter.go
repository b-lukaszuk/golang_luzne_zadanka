// simple solution,
// not entirely idiot/error-proof, not optimized for speed
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func getRandNumber(minIncl float64, maxExcl float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(maxExcl-minIncl) + minIncl
}

func celsiusToFahrenheit(cels float64) (float64, error) {
	if cels < -273.15 {
		return 0.0, errors.New("temp in deg cels musn't be below -273.15 deg cels (0 deg kelvin)")
	}
	return cels*(9.0/5.0) + 32.0, nil
}

func kelvinToFahrenheit(kelv float64) (float64, error) {
	if kelv < 0 {
		return 0.0, errors.New("temp in deg kelvin musn't be negative")
	}
	return kelv*(9.0/5.0) - 459.67, nil
}

type conversion func(float64) (float64, error)

func printConversionExamples(
	conv conversion,
	fromUnits string, toUnits string,
	nExamples uint) {
	for i := 0; i < int(nExamples); i++ {
		temp := getRandNumber(-1000.0, 1000.0)
		tConv, err := conv(temp)
		if err == nil {
			fmt.Printf("%.2f %s is %.2f %s\n", temp, fromUnits, tConv, toUnits)
		} else {
			fmt.Printf("Cannot convert %.2f %s to %s, %s\n",
				temp, fromUnits, toUnits, err)
		}
	}
}

func main() {
	fmt.Printf("Toy program. Temperature Converter.\n")
	fmt.Printf("No guarantee the results will be correct.\n\n")
	printConversionExamples(celsiusToFahrenheit, "celsius", "fahrenheit", 5)
	printConversionExamples(kelvinToFahrenheit, "kelvin", "fahrenheit", 5)
	fmt.Printf("\nThat's all. Goodbye!\n")
}
