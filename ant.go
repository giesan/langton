package main

import (
	"fmt"
)

// These are constants for the movement directions.
const (
	N int = iota // North
	E            // East
	S            // South
	W            // West
)

// Ant holds all propoeties of ant.
type Ant struct {
	Direction int
	Symbol    string
	Position
}

// Position holds the position coordinates.
type Position struct {
	Row, Column int
}

// Init initialises the ant with the symbol and the initial direction of movement.
func (a *Ant) Init() {
	a.Symbol = "X"
	a.Direction = N
}

// moveAnt changes the color at the current position and move the ant to the next position.
// Returns error if the end of the board is reached.
func moveAnt(board *Board, ant *Ant) error {
	toggleColor(board, ant.Position)
	setNextDirection(board, ant)
	setNextPosition(ant)
	if (*ant).Position.Row > board.Size-1 || (*ant).Position.Row < 0 ||
		(*ant).Position.Column > board.Size-1 || (*ant).Position.Column < 0 {
		return fmt.Errorf("end of the board reached - board size(%v)", board.Size)
	}
	return nil
}

// setNextDirection determines the direction for the next step and sets its.
func setNextDirection(board *Board, ant *Ant) {
	switch (*board).Matrix[ant.Position.Row][ant.Position.Column] {
	case white:
		if ant.Direction == W {
			(*ant).Direction = N
		} else {
			(*ant).Direction++
		}
	case black:
		if ant.Direction == N {
			(*ant).Direction = W
		} else {
			(*ant).Direction--
		}
	}
}

// setNextPosition determines the next position of the ant and set its.
func setNextPosition(ant *Ant) {
	switch ant.Direction {
	case N:
		(*ant).Position.Row--
	case E:
		(*ant).Position.Column++
	case S:
		(*ant).Position.Row++
	case W:
		(*ant).Position.Column--
	}
}

// posittionAntOnBoard sets the ant into center of the board.
// Returns error if the size of board is not correct.
func positionAntOnBoard(ant *Ant, size int) error {
	if size%2 == 1 {
		(*ant).Position.Row = size/2 + 1
		(*ant).Position.Column = size/2 + 1
	} else {
		return fmt.Errorf("board size is not correct: %v", size)
	}
	return nil
}
