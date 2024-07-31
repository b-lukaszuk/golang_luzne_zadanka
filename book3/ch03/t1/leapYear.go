package main

import (
	"fmt"
)

// literal translation of wikipedia's:
// "This extra leap day occurs in each year that is a multiple of 4, except for
// years evenly divisible by 100 but not by 400."

func isXDivisibleByN(x uint, n uint) bool {
	return x%n == 0
}

func isDivBy4(x uint) bool {
	return isXDivisibleByN(x, 4)
}

func isDivBy100butNot400(x uint) bool {
	return isXDivisibleByN(x, 100) && !isXDivisibleByN(x, 400)
}

func isLeapYear(yr uint) bool {
	var result bool = false
	if isDivBy4(yr) {
		result = true
		if isDivBy100butNot400(yr) {
			result = false
		}
	}
	return result
}

func main() {
	yrs := [...]uint{1700, 1800, 1900, 1904, 1914, 1950,
		1980, 1985, 2000, 2004, 2024}
	fmt.Println("Testing possible leap years.")
	for _, yr := range yrs {
		fmt.Printf("\tIs %d a leap year? %t\n", yr, isLeapYear(yr))
	}
	fmt.Println("That's all. Goodbye!")
}
