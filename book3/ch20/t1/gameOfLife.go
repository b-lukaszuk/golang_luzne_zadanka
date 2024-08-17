// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	width       = 80
	height      = 15
	probAlive   = 0.25
	aliveSymbol = "O"
	deadSymbol  = "."
	delayMs     = 500
	numCycles   = 50
)

type Universe [][]bool

func getNewUniverse() Universe {
	universe := make([][]bool, height)
	for r := 0; r < height; r++ {
		universe[r] = make([]bool, width)
	}
	return universe
}

func getRandNumber(minIncl float64, maxExcl float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(maxExcl-minIncl) + minIncl
}

func getRandAliveOrDead(probAlive float64) bool {
	var choice = getRandNumber(0, 1)
	if choice <= probAlive {
		return true
	}
	return false
}

func (u Universe) getRandState() {
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			u[r][c] = getRandAliveOrDead(probAlive)
		}
	}
}

// string representation of a field/cell in the universe
func getFieldSymbol(field bool) string {
	if field {
		return aliveSymbol
	}
	return deadSymbol
}

func (u Universe) print(cycleNum int) {
	var population uint = 0
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			fmt.Printf("%s", getFieldSymbol(u[r][c]))
			if u[r][c] {
				population++
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("cycle: %d, population: %d\n", cycleNum, population)
}

// the terminal must support ANSI escape codes
// https://en.wikipedia.org/wiki/ANSI_escape_code
func (u Universe) cleanPrintout() {
	// "\033[xxxxA" - xxx moves cursor up xxxx lines
	// +1 cause cycle/population line
	fmt.Printf("\033[" + strconv.Itoa(height+1) + "A")
	fmt.Printf("\033[J") // clears from cursor position till end of display
}

func (u Universe) reprint(cycleNum int) {
	u.cleanPrintout()
	u.print(cycleNum)
}

func isWithinRange(num, minInc, maxExcl int) bool {
	return !(num < minInc || num >= maxExcl)
}

func isCellWithinRange(x, y int) bool {
	return isWithinRange(x, 0, height) && isWithinRange(y, 0, width)
}

func (u Universe) getNumLiveNeighbours(x, y int) int {
	if !isCellWithinRange(x, y) {
		return 0
	}
	var count int = 0
	var xi, yi int = 0, 0
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			xi, yi = x+r, y+c
			if (xi == x && yi == y) || !isCellWithinRange(xi, yi) {
				continue
			}
			if u[xi][yi] {
				count++
			}
		}
	}
	return count
}

func (u Universe) shouldCellBeAlive(x, y int) bool {
	if u[x][y] {
		return isWithinRange(u.getNumLiveNeighbours(x, y), 2, 4)
	}
	return u.getNumLiveNeighbours(x, y) == 3
}

func (u Universe) getNextState() Universe {
	newUniverse := getNewUniverse()
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			newUniverse[r][c] = u.shouldCellBeAlive(r, c)
		}
	}
	return newUniverse
}

// prints the universe, runs cycle, repeats
func (u Universe) runGameOfLife(noOfCycles uint) {
	// initialization
	u = getNewUniverse()
	u.getRandState()
	u.print(0)

	// state transition
	for i := 0; i < int(noOfCycles); i++ {
		u = u.getNextState()
		u.reprint(i + 1) // disply cycle in human format 1-n (incl-incl)
		time.Sleep(delayMs * time.Millisecond)
	}
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It displays a so called game of life.\n")
	fmt.Printf("Note: your terminal must support ANSI escape codes.\n\n")

	usrInput := "" // start game of life animation on keypress
	fmt.Printf("Press Enter to begin.\n")
	fmt.Printf("You can abort at any time by pressing Ctrl+C.\n")
	fmt.Scanln(&usrInput)

	var universe Universe
	universe.runGameOfLife(numCycles)

	fmt.Printf("\nThat's all. Goodbye!\n")
}
