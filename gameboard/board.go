package gameboard

// Board holds all properties of board.
type Board struct {
	Size   int
	Matrix [][]string
}

// Position holds the position coordinates.
type Position struct {
	Row, Column int
}

// New creates and returns a new instance of Board with default values.
//
// The function initializes a new Board struct with an empty matrix and a size of 0.
// It does not allocate any memory for the matrix, as it will be initialized later when the Init method is called.
//
// Returns:
// - A new instance of Board with default values.
func New(size int) Board {
	var b Board
	b.Size = size
	b.Matrix = make([][]string, size)
	for i := range b.Matrix {
		b.Matrix[i] = make([]string, size)
	}
	b.Fill(White)
	return b
}

// Fill fills the entire board with the specified color.
//
// The function iterates over each cell in the board's matrix and assigns the given color to it.
// This operation is performed in-place, meaning that the original matrix is modified.
//
// Parameters:
// - color: A string representing the color to Fill the board with.
//
// Returns:
// This function does not return any value. The board's matrix is modified directly.
func (b *Board) Fill(color string) {
	for y, rows := range b.Matrix {
		for x := range rows {
			b.Matrix[y][x] = color
		}
	}
}
