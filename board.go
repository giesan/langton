package main

import "fmt"

// Board holds all proporties of board.
type Board struct {
	Size   int
	Matrix [][]string
}

// iniBoard creates and fills the board with base color
func initBoard(size int) Board {
	var board Board
	board.Size = size
	board.Matrix = make([][]string, size)
	for i := range board.Matrix {
		board.Matrix[i] = make([]string, size)
	}
	fillBoard(&board.Matrix, white)
	return board
}

// fillBoard fills the board with color
func fillBoard(board *[][]string, color string) {
	for y, rows := range *board {
		for x := range rows {
			(*board)[y][x] = color
		}
	}
}

// printBoard prints the board and ant to the standard out
func printBoard(board *Board, ant *Ant, step int) {
	fmt.Print("\033[H\033[2J")
	// fmt.Printf("\x1bc")
	// fmt.Printf("\x1b[2J")
	for y, row := range board.Matrix {
		for x, col := range row {
			if ant.Position.Row == y && ant.Position.Column == x {
				fmt.Printf(" %s", ant.Symbol)
			} else {
				fmt.Printf(" %s", col)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\nStep %+v\n", step)
}
