// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	width       = 9
	height      = 9
	empty       = 0
	path2Sudoku = "./sudoku_hard.txt"
)

func getWords(path string) ([]string, error) {
	// content - byteslice, defer f.Close() is inside ReadFile
	content, err := os.ReadFile(path)
	if err != nil {
		return make([]string, empty), err
	}
	text := string(content)         // byte slice to string
	textArr := strings.Fields(text) // split on whitespaces with Fields
	return textArr, nil
}

func putTextArr2Sudoku(textArr *[]string, sudoku *[height][width]uint64) {
	counter := 0
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			// Base 10, 64-bit size
			i, err := strconv.ParseUint((*textArr)[counter], 10, 64)
			if err == nil {
				sudoku[r][c] = i
			}
			counter++
		}
	}
}

func getSudoku(path string) ([height][width]uint64, error) {
	textArr, err := getWords(path)
	var sudoku [height][width]uint64
	if err != nil {
		return sudoku, err
	}
	putTextArr2Sudoku(&textArr, &sudoku)
	return sudoku, nil
}

func printSudoku(sudoku *[height][width]uint64) {
	// *3 to account for spaces btw nums, +4 to account for | seps
	fmt.Printf("%s\n", strings.Repeat("-", width*3+4))
	var row string = ""
	for r := 0; r < height; r++ {
		row = "|"
		for c := 0; c < width; c++ {
			row += fmt.Sprintf("% d ", sudoku[r][c])
			if (c+1)%3 == 0 {
				row += "|"
			}
		}
		fmt.Printf("%s\n", row)
		if (r+1)%3 == 0 {
			// *3 to account for spaces btw nums, +4 to account for | seps
			fmt.Printf("%s\n", strings.Repeat("-", width*3+4))
		}
	}
}

func isCollisionInCol(num uint64, curRow, curCol uint,
	sudoku *[height][width]uint64) bool {
	for r := 0; r < height; r++ {
		if uint(r) != curRow && sudoku[r][curCol] == num {
			return true
		}
	}
	return false
}

func isCollisionInRow(num uint64, curRow, curCol uint,
	sudoku *[height][width]uint64) bool {
	for c := 0; c < width; c++ {
		if uint(c) != curCol && sudoku[curRow][c] == num {
			return true
		}
	}
	return false
}

func isCollisionIn3x3(num uint64, curRow, curCol uint,
	sudoku *[height][width]uint64) bool {
	r0 := uint(math.Floor(float64(curRow)/3) * 3)
	c0 := uint(math.Floor(float64(curCol)/3) * 3)
	var iRow, iCol uint = 0, 0

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			iRow, iCol = r0+uint(r), c0+uint(c)
			if curRow == iRow && curCol == iCol {
				continue
			}
			if sudoku[iRow][iCol] == num {
				return true
			}
		}
	}
	return false
}

// is there a collision after placing a num in the field
func isCollision(num uint64, curRow, curCol uint,
	sudoku *[height][width]uint64) bool {
	return isCollisionIn3x3(num, curRow, curCol, sudoku) ||
		isCollisionInRow(num, curRow, curCol, sudoku) ||
		isCollisionInCol(num, curRow, curCol, sudoku)
}

func isSolved(sudoku *[height][width]uint64) bool {
	var sum uint64 = 0
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if sudoku[r][c] == empty {
				return false
			}
			sum += sudoku[r][c]
		}
	}
	return (sum == 405) // cause sum(1:9)*9, collision tests prevent duplicates
}

// modifies sudoku inplace, prints solution
func solve(sudoku *[height][width]uint64) {
	for r := uint(0); r < height; r++ {
		for c := uint(0); c < width; c++ {
			if sudoku[r][c] != empty {
				continue
			}
			for i := uint64(1); i < 10; i++ {
				if isCollision(i, r, c, sudoku) {
					continue
				}
				// tryout a value into a cell
				sudoku[r][c] = i
				if isSolved(sudoku) {
					// for solved sudoku
					printSudoku(sudoku)
					fmt.Println()
					return
				}
				// try to solve with that i value (see above) in
				solve(sudoku)
				// if that failed, assign empty again and try different i
				sudoku[r][c] = empty
			}
			return // if i = 1:9, does not solve then terminate
		}
	}
	return // terminate after passing all fields, when no solution found
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It solves a sudoku puzzle (if its solvable).\n\n")

	sudoku, err := getSudoku(path2Sudoku)
	if err == nil {
		printSudoku(&sudoku)
		fmt.Printf("\n")
		fmt.Printf("Solution(s) (empty line if none found):\n\n")
		solve(&sudoku)
	}

	fmt.Printf("\nThat's all. Goodbye!\n")
}
