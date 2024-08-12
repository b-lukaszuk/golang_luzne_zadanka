// simple solution,
// not entirely idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
)

// in Go arrays are passed by value, not by reference
func addPiecesToRow(board [8][8]string, pieces string, row uint) [8][8]string {
	for i, piece := range pieces {
		board[row][i] = string(piece)
	}
	return board
}

func addFigures(board [8][8]string) [8][8]string {
	board = addPiecesToRow(board, "RNBQKBNR", 0)
	board = addPiecesToRow(board, "rnbqkbnr", 7)
	return board
}

func addPawns(board [8][8]string) [8][8]string {
	board = addPiecesToRow(board, "PPPPPPPP", 1)
	board = addPiecesToRow(board, "pppppppp", 6)
	return board
}

func getSquare(symbol string, black bool) string {
	var square string = ""
	var fill string = " "
	if black {
		fill = "."
	}
	if symbol == "" {
		symbol = fill
	}
	square = fmt.Sprintf("%s%s%s", fill, symbol, fill)
	return square
}

func getRow(row [8]string, black bool) string {
	var rowStr string = ""
	for c := 0; c < len(row); c++ {
		rowStr += fmt.Sprintf("| %s ", getSquare(row[c], black))
		if c == (len(row) - 1) {
			rowStr += "|"
		}
		black = !black
	}
	return rowStr + "\n"
}

func getBoard(board [8][8]string) string {
	var boardStr string = ""
	var startBlack bool = false
	for r := len(board) - 1; r >= 0; r-- {
		boardStr += "-------------------------------------------------\n"
		boardStr += getRow(board[r], startBlack)
		startBlack = !startBlack
	}
	boardStr += "-------------------------------------------------\n"
	return boardStr
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It displays a chessboard (starting position).\n\n")

	var chessboard [8][8]string
	// in Go arrays are passed by value, not by reference
	chessboard = addFigures(chessboard)
	chessboard = addPawns(chessboard)

	fmt.Printf("%s\n", getBoard(chessboard))

	fmt.Printf("\nThat's all. Goodbye!\n")
}
