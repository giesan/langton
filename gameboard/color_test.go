package gameboard_test

import (
	"giesan/langton/gameboard"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToggleColor(t *testing.T) {
	tests := []struct {
		name          string
		initialColor  string
		expectedColor string
	}{
		{
			name:          "ToggleWhiteToBlack",
			initialColor:  gameboard.White,
			expectedColor: gameboard.Black,
		},
		{
			name:          "ToggleBlackToWhite",
			initialColor:  gameboard.Black,
			expectedColor: gameboard.White,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := gameboard.New(5)
			pos := gameboard.Position{Row: 2, Column: 2}
			board.Matrix[pos.Row][pos.Column] = tt.initialColor

			board.ToggleColor(pos)

			assert.Equal(t, tt.expectedColor, board.Matrix[pos.Row][pos.Column])
		})
	}
}
