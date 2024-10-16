package gameboard

// These are constants for the colors.
const (
	White = " "
	Black = "█"
)

// ToggleColor toggles the color of the position on the board.
//
// The function takes a pointer to a Board struct and a Position struct as parameters.
// The Board struct contains a 2D slice of strings representing the board's cells,
// where each string can be either " " (White) or "█" (Black).
// The Position struct contains the row and column indices of the position to be toggled.
//
// The function checks the color of the position on the board. If the color is Black,
// it sets the color to White. If the color is White, it sets the color to Black.
//
// The function does not return any value.
func (b *Board) ToggleColor(pos Position) {
	if (*b).Matrix[pos.Row][pos.Column] == Black {
		(*b).Matrix[pos.Row][pos.Column] = White
	} else {
		(*b).Matrix[pos.Row][pos.Column] = Black
	}
}
