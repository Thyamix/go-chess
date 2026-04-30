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

type TestPiece struct {
	piece Piece
	x     int
	y     int
}

func MakeTestBoard(pieces []TestPiece, threats [][2]int) Board {
	board := NewEmptyTestBoard()
	for i := range pieces {
		piece := pieces[i]
		board.SetPiece(piece.piece, piece.x, piece.y)
	}
	for i := range threats {
		coords := threats[i]
		board.threats[coords[1]] |= 0b00000001 << coords[0]
	}
	return board
}

func TestGetThreats(t *testing.T) {
	tests := map[string]struct {
		pieces          []TestPiece
		isBlackTurn     bool
		expectedThreats [][2]int
	}{
		"Single Pawn": {
			[]TestPiece{{PAWN.Black(), 4, 4}}, true, [][2]int{{3, 3}, {5, 3}}},
		"Single Pawn Edge": {
			[]TestPiece{{PAWN.White(), 7, 3}}, false, [][2]int{{6, 4}}},
		"Single Pawn Edge 2": {
			[]TestPiece{{PAWN.White(), 6, 3}}, false, [][2]int{{5, 4}, {7, 4}}},
		"Single Pawn Edge Top": {
			[]TestPiece{{PAWN.White(), 6, 7}}, false, [][2]int{}},
		"Single Pawn Edge Bottom": {
			[]TestPiece{{PAWN.Black(), 3, 0}}, true, [][2]int{}},
		"Double Pawn": {
			[]TestPiece{{PAWN.Black(), 4, 4}, {PAWN.Black(), 5, 4}}, true, [][2]int{{3, 3}, {4, 3}, {5, 3}, {6, 3}}},
		"Double Pawn with Attack": {
			[]TestPiece{{PAWN.Black(), 4, 4}, {PAWN.Black(), 5, 4}, {PAWN.White(), 5, 3}}, true, [][2]int{{3, 3}, {4, 3}, {5, 3}, {6, 3}}},
		"Single Knight Corner": {
			[]TestPiece{{KNIGHT.White(), 1, 0}}, false, [][2]int{{0, 2}, {2, 2}, {3, 1}}},
		"Single Knight Middle": {
			[]TestPiece{{KNIGHT.White(), 5, 5}}, false, [][2]int{{4, 3}, {6, 3}, {3, 4}, {3, 6}, {4, 7}, {6, 7}, {7, 6}, {7, 4}}},
		"Single King Middle": {
			[]TestPiece{{KING.White(), 5, 5}}, false, [][2]int{{4, 5}, {6, 5}, {6, 6}, {5, 6}, {4, 6}, {6, 4}, {5, 4}, {4, 4}}},
		"Single King Top": {
			[]TestPiece{{KING.White(), 5, 7}}, false, [][2]int{{4, 7}, {6, 7}, {4, 6}, {5, 6}, {6, 6}}},
		"Single King Bottom Corner": {
			[]TestPiece{{KING.White(), 0, 0}}, false, [][2]int{{0, 1}, {1, 1}, {1, 0}}},
		"Single Bishop Bottom Left Corner": {
			[]TestPiece{{BISHOP.Black(), 0, 0}}, true, [][2]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}}},
		"Single Bishop Top Right Corner": {
			[]TestPiece{{BISHOP.Black(), 7, 7}}, true, [][2]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {0, 0}}},
		"Single Bishop Center": {
			[]TestPiece{{BISHOP.Black(), 4, 4}}, true, [][2]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {5, 5}, {6, 6}, {7, 7},
				{5, 3}, {6, 2}, {7, 1}, {2, 6}, {1, 7}, {2, 6}, {3, 5}}},
		"Single Bishop Center Obstructed": {
			[]TestPiece{{BISHOP.White(), 4, 4}, {ROOK.Black(), 6, 6}, {PAWN.Black(), 6, 2}}, false, [][2]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {5, 5}, {6, 6}, {5, 3}, {6, 2}, {1, 7}, {2, 6}, {2, 6}, {3, 5}}},
		"Single Rook Center": {
			[]TestPiece{{ROOK.Black(), 3, 3}}, true, [][2]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {3, 5}, {3, 6}, {3, 7}, {0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 3}, {6, 3}, {7, 3}}},
		"Single Rook Center Obstructed": {
			[]TestPiece{{ROOK.Black(), 3, 3}, {KNIGHT.White(), 3, 5}, {KNIGHT.White(), 6, 3}}, true, [][2]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {3, 5}, {0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 3}, {6, 3}}},
	}

	for test := range tests {
		t.Run(test, func(t *testing.T) {
			expected := MakeTestBoard(tests[test].pieces, tests[test].expectedThreats)
			got := MakeTestBoard(tests[test].pieces, [][2]int{})

			got.GetThreats(tests[test].isBlackTurn)

			if got.threats != expected.threats {
				t.Errorf("got %v but wanted %v", got.threats, expected.threats)
			}
		})
	}
}
