package main

import (
	"fmt"
	"math/rand"
	"time"
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

func getRandNum(minIncl uint, maxExcl uint) uint {
	rand.Seed(time.Now().UnixNano())
	return uint(rand.Intn(int(maxExcl-minIncl)) + int(minIncl))
}

func getRandADYear() uint {
	return getRandNum(1, 4001)
}

func getRandMonth() uint {
	return getRandNum(1, 13)
}

func getRandDay(maxDayIncl uint) uint {
	return getRandNum(1, maxDayIncl+1)
}

func getRandDateAD() string {
	var era = "AD"
	year := getRandADYear()
	month := getRandMonth()
	daysInMonth := 31

	switch month {
	case 2:
		if isLeapYear(year) {
			daysInMonth = 29
		} else {
			daysInMonth = 28
		}
	case 4, 6, 9, 11:
		daysInMonth = 30
	}

	day := getRandDay(uint(daysInMonth))

	return fmt.Sprintf("y: %d %s, m: %d, d: %d", year, era, month, day)
}

func main() {
	fmt.Printf("Printing 10 random dates\n\n")
	for i := 0; i < 10; i++ {
		fmt.Println(getRandDateAD())
	}
	fmt.Printf("\nThat's all. Goodbye!\n")
}
