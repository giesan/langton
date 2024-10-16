package insect_test

import (
	"giesan/langton/gameboard"
	"giesan/langton/insect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name        string
		expectedSym string
		expectedDir int
		expectedRow int
		expectedCol int
	}{
		{
			name:        "NewAntInitialization",
			expectedSym: "X",
			expectedDir: 0, // Default value for int
			expectedRow: 0, // Default value for int
			expectedCol: 0, // Default value for int
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ant := insect.New()
			assert.Equal(t, tt.expectedSym, ant.Symbol)
			assert.Equal(t, tt.expectedDir, ant.Direction)
			assert.Equal(t, tt.expectedRow, ant.Position.Row)
			assert.Equal(t, tt.expectedCol, ant.Position.Column)
		})
	}
}

func TestInit(t *testing.T) {
	tests := []struct {
		name        string
		initialDir  int
		expectedDir int
	}{
		{
			name:        "InitializeNorth",
			initialDir:  insect.North,
			expectedDir: insect.North,
		},
		{
			name:        "InitializeEast",
			initialDir:  insect.East,
			expectedDir: insect.East,
		},
		{
			name:        "InitializeSouth",
			initialDir:  insect.South,
			expectedDir: insect.South,
		},
		{
			name:        "InitializeWest",
			initialDir:  insect.West,
			expectedDir: insect.West,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ant := insect.New()
			ant.Init(tt.initialDir)
			assert.Equal(t, tt.expectedDir, ant.Direction)
		})
	}
}

func TestMove(t *testing.T) {
	tests := []struct {
		name             string
		initialRow       int
		initialCol       int
		initialDir       int
		initialColor     string
		expectedRow      int
		expectedCol      int
		expectedDir      int
		expectedColor    string
		expectedError    bool
		expectedErrorMsg string
	}{
		{
			name:             "MoveWithinBoard",
			initialRow:       2,
			initialCol:       2,
			initialDir:       insect.North,
			initialColor:     gameboard.White,
			expectedRow:      2,
			expectedCol:      3,
			expectedDir:      insect.East,
			expectedColor:    gameboard.Black,
			expectedError:    false,
			expectedErrorMsg: "",
		},
		{
			name:             "MoveToEndOfBoard",
			initialRow:       2,
			initialCol:       4,
			initialDir:       insect.North,
			initialColor:     gameboard.White,
			expectedError:    true,
			expectedErrorMsg: "end of the board reached - board size(5)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := gameboard.New(5)
			board.Matrix[tt.initialRow][tt.initialCol] = tt.initialColor

			ant := insect.New()
			ant.Position.Row = tt.initialRow
			ant.Position.Column = tt.initialCol
			ant.Direction = tt.initialDir

			err := ant.Move(&board)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrorMsg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedRow, ant.Position.Row)
				assert.Equal(t, tt.expectedCol, ant.Position.Column)
				assert.Equal(t, tt.expectedDir, ant.Direction)
				assert.Equal(t, tt.expectedColor, board.Matrix[tt.initialRow][tt.initialCol])
			}
		})
	}
}

func TestNextDirection(t *testing.T) {
	tests := []struct {
		name        string
		initialRow  int
		initialCol  int
		initialDir  int
		cellColor   string
		expectedDir int
	}{
		{
			name:        "WhiteCellFacingNorth",
			initialRow:  2,
			initialCol:  2,
			initialDir:  insect.North,
			cellColor:   gameboard.White,
			expectedDir: insect.East,
		},
		{
			name:        "WhiteCellFacingEast",
			initialRow:  2,
			initialCol:  2,
			initialDir:  insect.East,
			cellColor:   gameboard.White,
			expectedDir: insect.South,
		},
		{
			name:        "WhiteCellFacingSouth",
			initialRow:  2,
			initialCol:  2,
			initialDir:  insect.South,
			cellColor:   gameboard.White,
			expectedDir: insect.West,
		},
		{
			name:        "WhiteCellFacingWest",
			initialRow:  2,
			initialCol:  2,
			initialDir:  insect.West,
			cellColor:   gameboard.White,
			expectedDir: insect.North,
		},
		{
			name:        "BlackCellFacingNorth",
			initialRow:  2,
			initialCol:  2,
			initialDir:  insect.North,
			cellColor:   gameboard.Black,
			expectedDir: insect.West,
		},
		{
			name:        "BlackCellFacingEast",
			initialRow:  2,
			initialCol:  2,
			initialDir:  insect.East,
			cellColor:   gameboard.Black,
			expectedDir: insect.North,
		},
		{
			name:        "BlackCellFacingSouth",
			initialRow:  2,
			initialCol:  2,
			initialDir:  insect.South,
			cellColor:   gameboard.Black,
			expectedDir: insect.East,
		},
		{
			name:        "BlackCellFacingWest",
			initialRow:  2,
			initialCol:  2,
			initialDir:  insect.West,
			cellColor:   gameboard.Black,
			expectedDir: insect.South,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := gameboard.New(5)
			board.Matrix[tt.initialRow][tt.initialCol] = tt.cellColor

			ant := insect.New()
			ant.Position.Row = tt.initialRow
			ant.Position.Column = tt.initialCol
			ant.Direction = tt.initialDir

			ant.NextDirection(&board)

			assert.Equal(t, tt.expectedDir, ant.Direction)
		})
	}
}

func TestNextPosition(t *testing.T) {
	tests := []struct {
		name        string
		initialRow  int
		initialCol  int
		direction   int
		expectedRow int
		expectedCol int
	}{
		{
			name:        "NorthDirection",
			initialRow:  2,
			initialCol:  2,
			direction:   insect.North,
			expectedRow: 1,
			expectedCol: 2,
		},
		{
			name:        "EastDirection",
			initialRow:  2,
			initialCol:  2,
			direction:   insect.East,
			expectedRow: 2,
			expectedCol: 3,
		},
		{
			name:        "SouthDirection",
			initialRow:  2,
			initialCol:  2,
			direction:   insect.South,
			expectedRow: 3,
			expectedCol: 2,
		},
		{
			name:        "WestDirection",
			initialRow:  2,
			initialCol:  2,
			direction:   insect.West,
			expectedRow: 2,
			expectedCol: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ant := insect.New()
			ant.Position.Row = tt.initialRow
			ant.Position.Column = tt.initialCol
			ant.Direction = tt.direction

			ant.NextPosition()

			assert.Equal(t, tt.expectedRow, ant.Position.Row)
			assert.Equal(t, tt.expectedCol, ant.Position.Column)
		})
	}
}

func TestPositionOnBoard(t *testing.T) {
	tests := []struct {
		name             string
		boardSize        int
		expectedRow      int
		expectedCol      int
		expectedError    bool
		expectedErrorMsg string
	}{
		{
			name:        "OddSize",
			boardSize:   5,
			expectedRow: 3,
			expectedCol: 3,
		},
		{
			name:             "EvenSize",
			boardSize:        4,
			expectedError:    true,
			expectedErrorMsg: "board size is not correct: 4",
		},
		{
			name:        "MinOddSize",
			boardSize:   1,
			expectedRow: 1,
			expectedCol: 1,
		},
		{
			name:        "LargerOddSize",
			boardSize:   7,
			expectedRow: 4,
			expectedCol: 4,
		},
		{
			name:             "ZeroSize",
			boardSize:        0,
			expectedError:    true,
			expectedErrorMsg: "board size is not correct: 0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := gameboard.New(tt.boardSize)
			ant := insect.New()
			err := ant.PositionOnBoard(&board)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrorMsg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedRow, ant.Position.Row)
				assert.Equal(t, tt.expectedCol, ant.Position.Column)
			}
		})
	}
}
