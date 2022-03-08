package main

import (
	"fmt"
	"os"
	"time"
)

const (
	boardSize = 11
)

func main() {
	var ant Ant
	ant.Init()

	board := initBoard(boardSize)

	err := positionAntOnBoard(&ant, board.Size)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var step int
	for {
		err := moveAnt(&board, &ant)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		step++
		printBoard(&board, &ant, step)
		time.Sleep(350 * time.Millisecond)
	}

}
