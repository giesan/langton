package gameboard_test

import (
	"giesan/langton/gameboard"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{
			name:     "SmallBoard",
			size:     3,
			expected: 3,
		},
		{
			name:     "LargeBoard",
			size:     10,
			expected: 10,
		},
		{
			name:     "MinimumBoard",
			size:     1,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := gameboard.New(tt.size)

			assert.Equal(t, tt.expected, board.Size)
			assert.Equal(t, tt.expected, len(board.Matrix))
			assert.Equal(t, tt.expected, len(board.Matrix[0]))

			for _, row := range board.Matrix {
				for _, cell := range row {
					assert.Equal(t, gameboard.White, cell)
				}
			}
		})
	}
}

func TestFill(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		color    string
		expected string
	}{
		{
			name:     "FillWithWhite",
			size:     3,
			color:    gameboard.White,
			expected: gameboard.White,
		},
		{
			name:     "FillWithBlack",
			size:     5,
			color:    gameboard.Black,
			expected: gameboard.Black,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := gameboard.New(tt.size)
			board.Fill(tt.color)

			for _, row := range board.Matrix {
				for _, cell := range row {
					assert.Equal(t, tt.expected, cell)
				}
			}
		})
	}
}
