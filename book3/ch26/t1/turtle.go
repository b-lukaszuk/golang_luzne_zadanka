// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	width        = 80
	height       = 15
	turtleSymbol = "Ŏ"
	emptySymbol  = "."
	delayMs      = 500
	numCycles    = 50
)

type gameField [][]bool

func getFieldSymbol(field bool) string {
	if field {
		return turtleSymbol
	}
	return emptySymbol
}

func getEmptyGameField() gameField {
	gameField := make([][]bool, height)
	for r := 0; r < height; r++ {
		gameField[r] = make([]bool, width)
	}
	return gameField
}

func (gf *gameField) setTurtle(rowNum, colNum uint) {
	(*gf)[rowNum][colNum] = true
}

func isWithinRange(num, minInc, maxExcl int) bool {
	return !(num < minInc || num >= maxExcl)
}

func isCellWithinRange(x, y int) bool {
	return isWithinRange(x, 0, height) && isWithinRange(y, 0, width)
}

func (gf *gameField) getTurtlePosition() (int, int) {
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if (*gf)[r][c] {
				return r, c
			}
		}
	}
	return 0, 0
}

func (gf *gameField) moveTurtle(horizontalShift, verticalShift int) {
	x, y := (*gf).getTurtlePosition()
	newX, newY := x+horizontalShift, y+verticalShift
	if !isCellWithinRange(newX, newY) {
		return
	}
	(*gf)[x][y] = false // remove turtle from old position
	(*gf).setTurtle(uint(newX), uint(newY))
}

func (gf *gameField) print() {
	turtlePos := ""
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			fmt.Printf("%s", getFieldSymbol((*gf)[r][c]))
			if (*gf)[r][c] {
				turtlePos = fmt.Sprintf("(%d, %d)", r, c)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("turtle position: %s.\n", turtlePos)
}

// the terminal must support ANSI escape codes
// https://en.wikipedia.org/wiki/ANSI_escape_code
func (gf *gameField) cleanPrintout() {
	// "\033[xxxxA" - xxx moves cursor up xxxx lines
	// +5 lines below gamefield itself
	fmt.Printf("\033[" + strconv.Itoa(height+5) + "A")
	fmt.Printf("\033[J") // clears from cursor position till end of display
}

func (gf *gameField) reprint() {
	(*gf).cleanPrintout()
	(*gf).print()
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It displays a turtle and allows you to move it.\n")
	fmt.Printf("Press Ctrl+C to abort the game.\n\n")

	gf := getEmptyGameField()
	gf.setTurtle(7, 40)
	gf.print()

	usrInput := "k"

	for strings.Contains("kjhl", usrInput) {
		fmt.Printf("\nk - UP, j - DOWN, h - LEFT, l - RIGHT.\n")
		fmt.Printf("q (or anything else) quits.\n")
		// acting on keypress requires non-standard libraries
		fmt.Printf("Your choice (click Enter to confirm): ")
		usrInput = ""
		fmt.Scanln(&usrInput)
		switch usrInput {
		case "k":
			gf.moveTurtle(-1, 0)
		case "j":
			gf.moveTurtle(1, 0)
		case "h":
			gf.moveTurtle(0, -1)
		case "l":
			gf.moveTurtle(0, 1)
		default:
			break
		}
		gf.reprint()
	}

	fmt.Printf("\nThat's all. Goodbye!\n")
}
