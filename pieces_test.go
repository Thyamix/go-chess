package gochess

import "testing"

func TestPieceGetType(t *testing.T) {
	tests := map[string]struct {
		piece    Piece
		expected Piece
	}{
		"Pawn - White":   {PAWN.White(), PAWN},
		"Knight - White": {KNIGHT.White(), KNIGHT},
		"Bishop - White": {BISHOP.White(), BISHOP},
		"Rook - White":   {ROOK.White(), ROOK},
		"Queen - White":  {QUEEN.White(), QUEEN},
		"King - White":   {KING.White(), KING},
		"Pawn - Black":   {PAWN.Black(), PAWN},
		"Knight - Black": {KNIGHT.Black(), KNIGHT},
		"Bishop - Black": {BISHOP.Black(), BISHOP},
		"Rook - Black":   {ROOK.Black(), ROOK},
		"Queen - Black":  {QUEEN.Black(), QUEEN},
		"King - Black":   {KING.Black(), KING},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.piece.Type() != test.expected {
				t.Errorf("got %v want %v", test.piece.Type(), test.expected)
			}
		})
	}
}
