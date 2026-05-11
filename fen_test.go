package gochess

import "testing"

func TestGetPosition(t *testing.T) {
	tests := map[string]struct {
		fen   FEN
		check func(t *testing.T, position *Position)
	}{
		// Standard starting positions

		//Board
		"starting position - white pieces": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.board[WHITE] != 0x000000000000FFFF {
					t.Errorf("white pieces: got %d, want 65535", position.board[WHITE])
				}
			},
		},
		"starting position - black pieces": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.board[BLACK] != 0xFF00000 {
					t.Errorf("black pieces: got %d, want 18446462598732840960", position.board[BLACK])
				}
			},
		},

		//Pawns
		"starting position - white pawns": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.pawns[WHITE] != 0x000000000000FF00 {
					t.Errorf("white pawns: got %d, want 65280", position.pawns[WHITE])
				}
			},
		},
		"starting position - black pawns": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.pawns[BLACK] != 0x00FF000000000000 {
					t.Errorf("black pawns: got %d, want 71776119061217280", position.pawns[BLACK])
				}
			},
		},

		// Knights
		"starting position - white knights": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.knight[WHITE] != 0x0000000000000042 {
					t.Errorf("white knight: got %d, want 66", position.knight[WHITE])
				}
			},
		},
		"starting position - black knights": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.knight[BLACK] != 0x4200000000000000 {
					t.Errorf("black knight: got %d, want 4755801206503243776", position.knight[BLACK])
				}
			},
		},
		// Bishops
		// Rooks
		// Kings
		// Queens
	}

	for test := range tests {
		t.Run(test, func(t *testing.T) {
			position := tests[test].fen.getPosition()
			tests[test].check(t, position)
		})
	}
}
