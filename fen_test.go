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
					t.Errorf("white pieces: got %016X, want 0x000000000000FFFF", position.board[WHITE])
				}
			},
		},
		"starting position - black pieces": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.board[BLACK] != 0xFFFF000000000000 {
					t.Errorf("black pieces: got %016X, want 0xFFFF000000000000", position.board[BLACK])
				}
			},
		},

		//Pawns
		"starting position - white pawns": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.pawns[WHITE] != 0x000000000000FF00 {
					t.Errorf("white pawns: got %016X, want 0x000000000000FF00", position.pawns[WHITE])
				}
			},
		},
		"starting position - black pawns": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.pawns[BLACK] != 0x00FF000000000000 {
					t.Errorf("black pawns: got %016X, want 0x00FF000000000000", position.pawns[BLACK])
				}
			},
		},

		// Knights
		"starting position - white knights": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.knights[WHITE] != 0x0000000000000042 {
					t.Errorf("white knight: got %016X, want 0x0000000000000042", position.knights[WHITE])
				}
			},
		},
		"starting position - black knights": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.knights[BLACK] != 0x4200000000000000 {
					t.Errorf("black knight: got %016X, want 0x4200000000000000", position.knights[BLACK])
				}
			},
		},

		// Bishops
		"starting position - white bishops": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.bishops[WHITE] != 0x0000000000000024 {
					t.Errorf("white bishops: got %016X, want 0x0000000000000024", position.bishops[WHITE])
				}
			},
		},
		"starting position - black bishops": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.bishops[BLACK] != 0x2400000000000000 {
					t.Errorf("black bishops: got %016X, want 0x2400000000000000", position.bishops[BLACK])
				}
			},
		},

		// Rooks
		"starting position - white rooks": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.rooks[WHITE] != 0x0000000000000081 {
					t.Errorf("white rooks: got %016X, want 0x0000000000000081", position.rooks[WHITE])
				}
			},
		},
		"starting position - black rooks": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.rooks[BLACK] != 0x8100000000000000 {
					t.Errorf("black rooks: got %016X, want 0x8100000000000000", position.rooks[BLACK])
				}
			},
		},

		// Queens
		"starting position - white queens": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.queens[WHITE] != 0x0000000000000010 {
					t.Errorf("white queens: got %016X, want 0x0000000000000010", position.queens[WHITE])
				}
			},
		},
		"starting position - black queens": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.queens[BLACK] != 0x1000000000000000 {
					t.Errorf("black queens: got %016X, want 0x1000000000000000", position.queens[BLACK])
				}
			},
		},

		// Kings
		"starting position - white kings": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.kings[WHITE] != 0x0000000000000008 {
					t.Errorf("white kings: got %016X, want 0x0000000000000008", position.kings[WHITE])
				}
			},
		},
		"starting position - black kings": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				if position.kings[BLACK] != 0x0800000000000000 {
					t.Errorf("black kings: got %016X, want 0x0800000000000000", position.kings[BLACK])
				}
			},
		},

		// Mailbox

		"starting position - mailbox": {
			fen: StartingPositionFEN,
			check: func(t *testing.T, position *Position) {
				expected := [64]Piece{13, 11, 11, 13, 15, 11, 11, 13, 9, 9, 9, 9, 9, 9, 9, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 4, 2, 3, 5, 6, 3, 2, 4}
				if position.pieces != expected {
					t.Errorf("mailbox: got %v, want %v", position.pieces, expected)
				}
			},
		},

		// Custom position

		// En passant

		// Castling rights

		// Move clock (half/full moves)

		// Side to move
	}

	for test := range tests {
		t.Run(test, func(t *testing.T) {
			position, err := tests[test].fen.getPosition()
			if err != nil {
				t.Error(err)
			}
			tests[test].check(t, position)
		})
	}
}
