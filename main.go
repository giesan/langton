package main

import (
	"fmt"

	"giesan/langton/gameboard"
	"giesan/langton/insect"
	"os"
	"time"
)

const (
	boardSize = 11
)

func main() {
	ant := insect.New()
	ant.Init(insect.South)

	board := gameboard.New(boardSize)

	err := ant.PositionOnBoard(&board)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	runAnt(&ant, &board)

}

// This function continuously moves the ant on the board according to the Langton's ant rules.
// It prints the board state after each move and waits for 100 milliseconds before the next move.
// The function runs indefinitely until an error occurs.
//
// Parameters:
// - a: A pointer to the ant object that will be moving on the board.
// - b: A pointer to the board object where the ant will move.
// - step: An integer representing the current step number. It is incremented after each move.
func runAnt(a *insect.Ant, b *gameboard.Board) {
	var step int
	for {
		err := a.Move(b)
		if err != nil {
			fmt.Println(err)
			return
		}
		step++
		printBoard(b, a, step)
		time.Sleep(100 * time.Millisecond)
	}
}
