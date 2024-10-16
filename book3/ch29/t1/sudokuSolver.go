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
	path2Sudoku = "./sudoku.txt"
)

func getWords(path string) ([]string, error) {
	// content - byteslice, defer f.Close() is inside ReadFile
	content, err := os.ReadFile(path)
	if err != nil {
		return make([]string, 0), err
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
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			fmt.Printf("%d\t", sudoku[r][c])
		}
		fmt.Printf("\n")
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

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It solves a sudoku puzzle (if its solvable).\n\n")

	sudoku, err := getSudoku(path2Sudoku)
	if err == nil {
		printSudoku(&sudoku)
	}

	fmt.Printf("\nThat's all. Goodbye!\n")
}
