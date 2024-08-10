// simple solution,
// not entirely idiot/error-proof, not optimized for speed
package main

import (
	"errors"
	"fmt"
)

type conversion func(float64) (float64, error)

func celsiusToFahrenheit(cels float64) (float64, error) {
	if cels < -273.15 {
		return 0.0, errors.New("temp in deg cels musn't be below -273.15 deg cels (0 deg kelvin)")
	}
	return cels*(9.0/5.0) + 32.0, nil
}

func fahrenheitToCelsius(fahr float64) (float64, error) {
	if fahr < -459.67 {
		return 0.0, errors.New("temp in deg fahr musn't be below -459.67 deg (0 deg kelvin)")
	}
	return (fahr - 32.0) * (5.0 / 9.0), nil
}

func getRow(temp float64, conv conversion) string {
	var temp2, err = conv(temp)
	if err != nil {
		return fmt.Sprintf("| %-8.2f | %-8s |\n", temp, "Err")
	}
	return fmt.Sprintf("| %-8.2f | %-8.2f |\n", temp, temp2)
}

func getHeader(fromUnit string, toUnit string) string {
	return fmt.Sprintf("| %-8s | %-8s |\n", fromUnit, toUnit) +
		"=======================\n"
}
func getFooter() string {
	return "=======================\n"
}

func getTable(conv conversion, fromUnit string, toUnit string) string {
	var table string = getHeader(fromUnit, toUnit)
	// by 20 to shorten the output
	for i := -40; i < 101; i += 20 {
		table += getRow(float64(i), conv)
	}
	return table + getFooter()
}

func main() {
	fmt.Printf("Toy program. Temperature Converter Tables.\n")
	fmt.Printf("No guarantee the output is correct.\n\n")

	fmt.Printf("%s\n", getTable(celsiusToFahrenheit, "deg. C", "deg. F"))
	fmt.Printf("%s", getTable(fahrenheitToCelsius, "deg. F", "deg. C"))

	fmt.Printf("\nThat's all. Goodbye!\n")
}
