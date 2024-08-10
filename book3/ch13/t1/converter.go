// simple solution,
// not entirely idiot/error-proof, not optimized for speed
package main

import (
	"errors"
	"fmt"
)

type celsius float64
type fahrenheit float64
type kelvin float64

func (c celsius) toFahrenheit() (fahrenheit, error) {
	if c < -273.15 {
		return fahrenheit(0.0), errors.New("temp in deg cels musn't be below -273.15 deg cels (0 deg kelvin)")
	}
	return fahrenheit(c*(9.0/5.0) + 32.0), nil
}

func (c celsius) toKelvin() (kelvin, error) {
	if c < -273.15 {
		return kelvin(0.0), errors.New("temp in deg cels musn't be below -273.15 deg cels (0 deg kelvin)")
	}
	return kelvin(c + 273.15), nil
}

func (k kelvin) toFahrenheit() (fahrenheit, error) {
	if k < 0 {
		return fahrenheit(0.0), errors.New("temp in deg kelv musn't be below 0 deg kelvin")
	}
	return fahrenheit(k*(9.0/5.0) - 459.67), nil
}

func (k kelvin) toCelsius() (celsius, error) {
	if k < 0 {
		return celsius(0.0), errors.New("temp in deg kelv musn't be below 0 deg kelvin")
	}
	return celsius(k - 273.15), nil
}

func (f fahrenheit) toKelvin() (kelvin, error) {
	if f < 0 {
		return kelvin(0.0), errors.New("temp in deg fahr musn't be below -459.67 deg (0 deg kelvin)")
	}
	return kelvin((f + 459.67) * (5.0 / 9.0)), nil
}

func (f fahrenheit) toCelsius() (celsius, error) {
	if f < 0 {
		return celsius(0.0), errors.New("temp in deg fahr musn't be below -459.67 deg (0 deg kelvin)")
	}
	return celsius((f - 32.0) * (5.0 / 9.0)), nil
}

func main() {
	fmt.Printf("Toy program. Temperature Converter.\n")
	fmt.Printf("No guarantee the results will be correct.\n\n")

	var c celsius = 100.0
	var f1, _ = c.toFahrenheit()
	var k1, _ = c.toKelvin()
	fmt.Printf("%.2f deg celsius is %.2f deg fahrenheit.\n", c, f1)
	fmt.Printf("%.2f deg celsius is %.2f deg kelvin.\n", c, k1)

	var f fahrenheit = 100.0
	var c2, _ = f.toCelsius()
	var k2, _ = f.toKelvin()
	fmt.Printf("%.2f deg fahrenheit is %.2f deg celsius.\n", f, c2)
	fmt.Printf("%.2f deg fahrenheit is %.2f deg kelvin.\n", f, k2)

	var k kelvin = 100.0
	var c3, _ = k.toCelsius()
	var f3, _ = k.toFahrenheit()
	fmt.Printf("%.2f deg kelvin is %.2f deg celsius.\n", k, c3)
	fmt.Printf("%.2f deg kelvin is %.2f deg fahrenheit.\n", k, f3)

	fmt.Printf("\nThat's all. Goodbye!\n")
}
