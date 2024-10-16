package main

import (
	"fmt"
	"giesan/langton/gameboard"
	"giesan/langton/insect"
)

// printBoard prints the board and ant to the standard out.
// It uses ANSI escape codes to clear the terminal screen before printing the new board state.
// The function iterates over the board matrix, printing the ant symbol at its current position and
// the board cell contents at other positions.
// It also prints the current step number at the end of the board.
//
// Parameters:
// - board: A pointer to the board struct containing the matrix of cells.
// - a: A pointer to the ant struct representing the ant's current position and direction.
// - step: An integer representing the current step number in the simulation.
func printBoard(board *gameboard.Board, a *insect.Ant, step int) {
	fmt.Print("\033[H\033[2J")
	// fmt.Printf("\x1bc")
	// fmt.Printf("\x1b[2J")
	for y, row := range board.Matrix {
		for x, col := range row {
			if a.Position.Row == y && a.Position.Column == x {
				fmt.Printf(" %s", a.Symbol)
			} else {
				fmt.Printf(" %s", col)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\nStep %+v\n", step)
}
