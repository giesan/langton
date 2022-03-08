package main

// These are constants for the colors.
const (
	white = " "
	black = "█"
)

// toggleColor toggles the color of the position.
func toggleColor(board *Board, pos Position) {
	if (*board).Matrix[pos.Row][pos.Column] == black {
		(*board).Matrix[pos.Row][pos.Column] = white
	} else {
		(*board).Matrix[pos.Row][pos.Column] = black
	}
}
