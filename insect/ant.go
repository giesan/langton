package insect

import (
	"fmt"
	"giesan/langton/gameboard"
)

// These are constants for the movement directions.
const (
	North int = iota // North
	East             // East
	South            // South
	West             // West
)

// Ant holds all properties of ant.
type Ant struct {
	Direction int
	Symbol    string
	Position  gameboard.Position
}

// New creates and initializes a new Ant instance with a default symbol "X".
// The initial direction of movement is not set in this function.
// Use the Init method to set the initial direction.
func New() Ant {
	return Ant{
		Symbol: "X",
	}
}

// Init initializes the ant with the symbol and the initial direction of movement.
//
// Parameters:
// d (int): The initial direction of movement for the ant. It should be one of the following constants:
// - North: 0
// - East:  1
// - South: 2
// - West:  3
//
// The function does not return any value.
func (a *Ant) Init(d int) {
	a.Direction = d
}

// Move changes the color at the current position and moves the ant to the next position.
// It returns an error if the end of the board is reached.
//
// Parameters:
// b (board.Board): A pointer to the board on which the ant is moving.
//
// Return:
// error: An error indicating whether the end of the board is reached or not.
// If the end of the board is reached, it returns an error with a message indicating the board size.
// If the end of the board is not reached, it returns nil.
func (a *Ant) Move(b *gameboard.Board) error {
	a.NextDirection(b)
	b.ToggleColor(a.Position)
	a.NextPosition()
	if (*a).Position.Row > b.Size-1 ||
		(*a).Position.Row < 0 ||
		(*a).Position.Column > b.Size-1 ||
		(*a).Position.Column < 0 {
		return fmt.Errorf("end of the board reached - board size(%v)", b.Size)
	}
	return nil
}

// NextDirection determines the direction for the next step and sets its value.
// The ant's direction is based on the current color of the board cell it is on.
// If the cell is white, the ant turns right (clockwise).
// If the cell is black, the ant turns left (counterclockwise).
//
// Parameters:
// b (board.Board): A pointer to the board on which the ant is moving.
//
// Return:
// None. The function modifies the ant's direction in place.
func (a *Ant) NextDirection(b *gameboard.Board) {
	switch (*b).Matrix[a.Position.Row][a.Position.Column] {
	case gameboard.White:
		if a.Direction == West {
			(*a).Direction = North
		} else {
			(*a).Direction++
		}
	case gameboard.Black:
		if a.Direction == North {
			(*a).Direction = West
		} else {
			(*a).Direction--
		}
	}
}

// NextPosition determines the next position of the ant and sets its value.
// The ant's position is updated based on its current direction.
//
// Parameters:
// ant (Ant): A pointer to the Ant instance for which the next position needs to be calculated.
//
// Return:
// None. The function modifies the ant's position in place.
func (a *Ant) NextPosition() {
	switch a.Direction {
	case North:
		(*a).Position.Row--
	case East:
		(*a).Position.Column++
	case South:
		(*a).Position.Row++
	case West:
		(*a).Position.Column--
	}
}

// PositionOnBoard sets the initial position of the ant on the board.
// The ant's position is determined based on the board size.
// If the board size is odd, the ant is placed in the center of the board.
// If the board size is even, an error is returned.
//
// Parameters:
// b (board.Board): A pointer to the board on which the ant is moving.
//
// Return:
// error: An error indicating whether the board size is correct or not.
// If the board size is even, it returns an error with a message indicating the board size.
// If the board size is odd, it returns nil.
func (a *Ant) PositionOnBoard(b *gameboard.Board) error {
	size := b.Size
	if size%2 == 1 {
		(*a).Position.Row = size/2 + 1
		(*a).Position.Column = size/2 + 1
	} else {
		return fmt.Errorf("board size is not correct: %v", size)
	}
	return nil
}
