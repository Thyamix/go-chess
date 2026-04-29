package gochess

import (
	"testing"
)

func NewEmptyTestBoard() Board {
	board := [8]uint32{
		0b00000000000000000000000000000000,
		0b00000000000000000000000000000000,
		0b00000000000000000000000000000000,
		0b00000000000000000000000000000000,
		0b00000000000000000000000000000000,
		0b00000000000000000000000000000000,
		0b00000000000000000000000000000000,
		0b00000000000000000000000000000000,
	}
	return Board{board: board}
}

func TestNewBoard(t *testing.T) {
	test := [64]Piece{
		ROOK.White(), KNIGHT.White(), BISHOP.White(), QUEEN.White(), KING.White(), BISHOP.White(), KNIGHT.White(), ROOK.White(),
		PAWN.White(), PAWN.White(), PAWN.White(), PAWN.White(), PAWN.White(), PAWN.White(), PAWN.White(), PAWN.White(),
		EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY,
		EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY,
		EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY,
		EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY,
		PAWN.Black(), PAWN.Black(), PAWN.Black(), PAWN.Black(), PAWN.Black(), PAWN.Black(), PAWN.Black(), PAWN.Black(),
		ROOK.Black(), KNIGHT.Black(), BISHOP.Black(), QUEEN.Black(), KING.Black(), BISHOP.Black(), KNIGHT.Black(), ROOK.Black(),
	}
	t.Run("New Board", func(t *testing.T) {
		board := NewBoard()
		for y := range 8 {
			for x := range 8 {
				piece, _ := board.GetPiece(x, y)
				if piece != test[y*8+x] {
					t.Errorf("got %v wanted %v on x:%v y:%v", piece, test[y*8+x], x, y)
				}
			}
		}
	})
}

func TestSetPiece(t *testing.T) {
	type TestPlacement struct {
		piece    Piece
		pos      [2]int
		expected bool
	}

	tests := map[string][]TestPlacement{
		"Single Black Pawn": {
			{PAWN.Black(), [2]int{2, 2}, true},
		},
		"Double Pawn": {
			{PAWN.White(), [2]int{2, 4}, true},
			{PAWN.Black(), [2]int{6, 1}, true},
		},
		"Single White Knight": {
			{KNIGHT.White(), [2]int{7, 0}, true},
		},
		"Single Black King": {
			{KING.Black(), [2]int{2, 4}, true},
		},
		"Replace Rook with Black Queen": {
			{ROOK.White(), [2]int{4, 4}, false},
			{QUEEN.Black(), [2]int{4, 4}, true},
		},
	}

	for test := range tests {
		t.Run(test, func(t *testing.T) {
			board := NewEmptyTestBoard()
			placements := tests[test]
			for i := range placements {
				placement := placements[i]
				board.SetPiece(placement.piece, placement.pos[0], placement.pos[1])
			}
			for i := range placements {
				placement := placements[i]
				piece, _ := board.GetPiece(placement.pos[0], placement.pos[1])
				if placement.expected {
					if piece != placement.piece {
						t.Errorf("want %v at %v,%v got %v", placement.piece, placement.pos[0], placement.pos[1], piece)
					}
				} else {
					if piece == placement.piece {
						t.Errorf("expected %v to be replaced at %v,%v", placement.piece, placement.pos[0], placement.pos[1])
					}
				}
			}
		})
	}
}

func TestGetThreats(t *testing.T) {

}
