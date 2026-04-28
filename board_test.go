package gochess

import (
	"testing"
)

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
				piece, _ := board.Get(x, y)
				if *piece != test[y*8+x] {
					t.Errorf("got %v wanted %v on x:%v y:%v", *piece, test[y*8+x], x, y)
				}
			}
		}
	})
}
