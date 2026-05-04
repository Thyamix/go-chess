package gochess

import "testing"

func TestPieceGetType(t *testing.T) {
	tests := map[string]struct {
		piece    Piece
		expected Piece
	}{
		"Pawn - White":   {PAWN.white(), PAWN},
		"Knight - White": {KNIGHT.white(), KNIGHT},
		"Bishop - White": {BISHOP.white(), BISHOP},
		"Rook - White":   {ROOK.white(), ROOK},
		"Queen - White":  {QUEEN.white(), QUEEN},
		"King - White":   {KING.white(), KING},
		"Pawn - Black":   {PAWN.black(), PAWN},
		"Knight - Black": {KNIGHT.black(), KNIGHT},
		"Bishop - Black": {BISHOP.black(), BISHOP},
		"Rook - Black":   {ROOK.black(), ROOK},
		"Queen - Black":  {QUEEN.black(), QUEEN},
		"King - Black":   {KING.black(), KING},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.piece.Type() != test.expected {
				t.Errorf("got %v want %v", test.piece.Type(), test.expected)
			}
		})
	}
}

func TestIsBlack(t *testing.T) {
	tests := []struct {
		piece    Piece
		expected bool
	}{
		{PAWN.black(), true},
		{KING.black(), true},
		{BISHOP.white(), false},
		{KNIGHT.white(), false},
	}

	for i := range tests {
		t.Run("IsBlack", func(t *testing.T) {
			if tests[i].piece.IsBlack() != tests[i].expected {
				t.Errorf("got %v expected %v", tests[i].piece.IsBlack(), tests[i].expected)
			}
		})
	}
}
