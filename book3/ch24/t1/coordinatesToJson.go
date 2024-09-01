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

// removed DMS field:
// looks bad on printout in terminal, repeats info from other fields
type coordinate struct {
	Decimal    float64
	Degrees    int
	Minutes    int
	Seconds    int
	Hemisphere string
}

func getAbs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// returns decimal degrees, minutes and seconds
func getDms(l location, takeLat bool) (float64, int, int, int) {
	var toConvert float64
	if takeLat {
		toConvert = l.Lat
	} else {
		toConvert = l.Long
	}
	var toConvertAbs float64 = getAbs(toConvert)
	var degs int = int(toConvertAbs)
	var mins float64 = (toConvertAbs - float64(degs)) * 60
	var secs float64 = (mins - float64(int(mins))) * 60
	return toConvert, degs, int(mins), int(secs)
}

func getSide(l location, takeLat bool) string {
	if takeLat {
		if l.Lat < 0 {
			return "S"
		} else {
			return "N"
		}
	} else {
		if l.Long < 0 {
			return "W"
		} else {
			return "E"
		}
	}
}

func getCoordinate(l location, takeLat bool) coordinate {
	var decim, degs, mins, secs = getDms(l, takeLat)
	var side string = getSide(l, takeLat)
	return coordinate{
		Decimal:    decim,
		Degrees:    degs,
		Minutes:    mins,
		Seconds:    secs,
		Hemisphere: side,
	}
}

func getCoordinates(l location) (coordinate, coordinate) {
	return getCoordinate(l, true), getCoordinate(l, false)
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
	fmt.Printf("It displays the (hopefully correct) coordinates as JSON.\n\n")

	spirit := location{-14.5684, 175.472636}
	c1, c2 := getCoordinates(spirit)
	fmt.Printf("The coordinates of spirit: %v.\n", spirit)
	bytes1, err1 := json.Marshal(c1)
	bytes2, err2 := json.Marshal(c2)
	exitOnError(err1)
	exitOnError(err2)

	fmt.Printf("Latitude: %s\n", string(bytes1))
	fmt.Printf("and\n")
	fmt.Printf("Longitude: %s\n", string(bytes2))

	fmt.Printf("\nThat's all. Goodbye!\n")
}
