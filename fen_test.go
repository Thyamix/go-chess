package gochess

import "testing"

func TestGetState(t *testing.T) {
	tests := map[string]struct {
		fen   FEN
		check func(t *testing.T, state *State, err error)
	}{
		// Standard starting state

		//Board
		"starting state - all pieces": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.board != 0xFFFF00000000FFFF {
					t.Errorf("white pieces: got %016X, want 0xFFFF00000000FFFF", state.white)
				}
			},
		},
		"starting state - white pieces": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.white != 0x000000000000FFFF {
					t.Errorf("white pieces: got %016X, want 0x000000000000FFFF", state.white)
				}
			},
		},
		"starting state - black pieces": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.black != 0xFFFF000000000000 {
					t.Errorf("black pieces: got %016X, want 0xFFFF000000000000", state.black)
				}
			},
		},

		//Pawns
		"starting state - white pawns": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.pawns[white] != 0x000000000000FF00 {
					t.Errorf("white pawns: got %016X, want 0x000000000000FF00", state.pawns[white])
				}
			},
		},
		"starting state - black pawns": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.pawns[black] != 0x00FF000000000000 {
					t.Errorf("black pawns: got %016X, want 0x00FF000000000000", state.pawns[black])
				}
			},
		},

		// Knights
		"starting state - white knights": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.knights[white] != 0x0000000000000042 {
					t.Errorf("white knight: got %016X, want 0x0000000000000042", state.knights[white])
				}
			},
		},
		"starting state - black knights": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.knights[black] != 0x4200000000000000 {
					t.Errorf("black knight: got %016X, want 0x4200000000000000", state.knights[black])
				}
			},
		},

		// Bishops
		"starting state - white bishops": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.bishops[white] != 0x0000000000000024 {
					t.Errorf("white bishops: got %016X, want 0x0000000000000024", state.bishops[white])
				}
			},
		},
		"starting state - black bishops": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.bishops[black] != 0x2400000000000000 {
					t.Errorf("black bishops: got %016X, want 0x2400000000000000", state.bishops[black])
				}
			},
		},

		// Rooks
		"starting state - white rooks": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.rooks[white] != 0x0000000000000081 {
					t.Errorf("white rooks: got %016X, want 0x0000000000000081", state.rooks[white])
				}
			},
		},
		"starting state - black rooks": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.rooks[black] != 0x8100000000000000 {
					t.Errorf("black rooks: got %016X, want 0x8100000000000000", state.rooks[black])
				}
			},
		},

		// Queens
		"starting state - white queens": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.queens[white] != 0x0000000000000010 {
					t.Errorf("white queens: got %016X, want 0x0000000000000010", state.queens[white])
				}
			},
		},
		"starting state - black queens": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.queens[black] != 0x1000000000000000 {
					t.Errorf("black queens: got %016X, want 0x1000000000000000", state.queens[black])
				}
			},
		},

		// Kings
		"starting state - white kings": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.kings[white] != 0x0000000000000008 {
					t.Errorf("white kings: got %016X, want 0x0000000000000008", state.kings[white])
				}
			},
		},
		"starting state - black kings": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.kings[black] != 0x0800000000000000 {
					t.Errorf("black kings: got %016X, want 0x0800000000000000", state.kings[black])
				}
			},
		},

		// Mailbox

		"starting state - mailbox": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				expected := [64]Piece{12, 10, 11, 13, 14, 11, 10, 12, 9, 9, 9, 9, 9, 9, 9, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 4, 2, 3, 5, 6, 3, 2, 4}
				if state.Pieces != expected {
					t.Errorf("mailbox: got %v, want %v", state.Pieces, expected)
				}
			},
		},

		// TODO: Custom state

		// Active colour
		"active colour - white": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.IsWhiteTurn {
					t.Errorf("isWhiteTurn: got %t, want false", state.IsWhiteTurn)
				}
			},
		},
		"active colour - black": {
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1",
			check: func(t *testing.T, state *State, err error) {
				if !state.IsWhiteTurn {
					t.Errorf("isWhiteTurn: got %t, want true", state.IsWhiteTurn)
				}
			},
		},
		"active colour - error": {
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR - KQkq - 0 1",
			check: func(t *testing.T, state *State, err error) {
				if err != ErrFenActiveColor {
					t.Errorf("isWhiteTurn: got no error, want error")
				}
			},
		},

		// Castling rights
		"castle right - default state": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.castle != 0b00001111 {
					t.Errorf("castle rights: got 0b%08b, want 0b00001111", state.castle)
				}
			},
		},
		"castle right - different rights": {
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w kQ - 0 1",
			check: func(t *testing.T, state *State, err error) {
				if state.castle != 0b00001001 {
					t.Errorf("castle rights: got 0b%08b, want 0b00001001", state.castle)
				}
			},
		},
		"castle right - empty": {
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w - - 0 1",
			check: func(t *testing.T, state *State, err error) {
				if state.castle != 0b00000000 {
					t.Errorf("castle rights: got 0b%08b, want 0b00001111", state.castle)
				}
			},
		},
		"castle right - invalid": {
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq- - 0 1",
			check: func(t *testing.T, state *State, err error) {
				if err != ErrFenCastleRights {
					t.Errorf("castle rights: should be invalid but got 0b%08b", state.castle)
				}
			},
		},
		"castle right - invalid chars": {
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w trsv - 0 1",
			check: func(t *testing.T, state *State, err error) {
				if err != ErrFenCastleRights {
					t.Errorf("castle rights: should be invalid but got 0b%08b", state.castle)
				}
			},
		},

		// En Passant Target
		"en passant target - empty": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.enPassant != 0 {
					t.Errorf("en passant target: got 0b%08b, want 0b00000000", state.enPassant)
				}
			},
		},
		"en passant target - e3": {
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq e3 0 1",
			check: func(t *testing.T, state *State, err error) {
				if state.enPassant != 0b00001000 {
					t.Errorf("en passant target: got 0b%08b, want 0b00001000", state.enPassant)
				}
			},
		},
		"en passant target - a6": {
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq a6 0 1",
			check: func(t *testing.T, state *State, err error) {
				if state.enPassant != 0b10000000 {
					t.Errorf("en passant target: got 0b%08b, want 0b10000000", state.enPassant)
				}
			},
		},
		"en passant target - h6": {
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq h6 0 1",
			check: func(t *testing.T, state *State, err error) {
				if state.enPassant != 0b00000001 {
					t.Errorf("en passant target: got 0b%08b, want 0b00000001", state.enPassant)
				}
			},
		},
		"en passant target - error": {
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq i6 0 1",
			check: func(t *testing.T, state *State, err error) {
				if err != ErrFenEnPassantTarget {
					t.Errorf("en passant target: got 0b%08b, want error", state.enPassant)
				}
			},
		},

		// Move clock (half/full moves)
		"half move clock - starting state": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.halfMove != 0 {
					t.Errorf("half move clock: got %v, want 0", state.halfMove)
				}
			},
		},
		"half move clock - 30": {
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 30 18",
			check: func(t *testing.T, state *State, err error) {
				if state.halfMove != 30 {
					t.Errorf("half move clock: got %v, want 0", state.halfMove)
				}
			},
		},
		"half move clock - 100": {
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 100 55",
			check: func(t *testing.T, state *State, err error) {
				if err != ErrFenHalfMoveTooHigh {
					t.Errorf("half move clock: got %v, want err", state.halfMove)
				}
			},
		},
		"full move clock - starting state": {
			fen: FENDefaultStart,
			check: func(t *testing.T, state *State, err error) {
				if state.Move != 1 {
					t.Errorf("move clock: got %v, want 1", state.Move)
				}
			},
		},
		"move clock - 30": {
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 30 30",
			check: func(t *testing.T, state *State, err error) {
				if state.Move != 30 {
					t.Errorf("half move clock: got %v, want 30", state.halfMove)
				}
			},
		},
		"move clock - 50": {
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 23 50",
			check: func(t *testing.T, state *State, err error) {
				if state.Move != 50 {
					t.Errorf("half move clock: got %v, want 50", state.halfMove)
				}
			},
		},
		"move > 2 x half move clock - 50": {
			fen: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 34 15",
			check: func(t *testing.T, state *State, err error) {
				if err != ErrFenHalfMoveToMoveRatio {
					t.Errorf("half move clock: got %v, want err", state.halfMove)
				}
			},
		},
	}

	for test := range tests {
		t.Run(test, func(t *testing.T) {
			state, err := tests[test].fen.getState()
			tests[test].check(t, state, err)
		})
	}
}
