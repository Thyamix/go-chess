package gochess

import "testing"

func TestAddPossibleMoves(t *testing.T) {
	tests := map[string]struct {
		pieces        []TestPiece
		isBlackTurn   bool
		expectedMoves []Move
	}{
		// All require king to check turn validity
		"Single Pawn":                   {[]TestPiece{{PAWN.black(), 4, 6}, {KING.black(), 7, 7}}, true, []Move{0x4645, 0x4644, 0x7766, 0x7776, 0x7767}},
		"Single Pawn Obstructed":        {[]TestPiece{{PAWN.black(), 4, 6}, {PAWN.white(), 4, 5}, {KING.black(), 7, 7}}, true, []Move{0x7766, 0x7776, 0x7767}},
		"Single Pawn Take":              {[]TestPiece{{PAWN.black(), 4, 6}, {PAWN.white(), 5, 5}, {KING.black(), 7, 7}}, true, []Move{0x4645, 0x4644, 0x4655, 0x7776, 0x7767}},
		"Many Pawns":                    {[]TestPiece{{PAWN.black(), 4, 6}, {PAWN.black(), 5, 6}, {PAWN.black(), 6, 6}, {PAWN.black(), 7, 4}, {PAWN.white(), 5, 5}, {KING.black(), 7, 7}}, true, []Move{0x7473, 0x6665, 0x6664, 0x6655, 0x4655, 0x4645, 0x4644, 0x7776, 0x7767}},
		"Single Knight":                 {[]TestPiece{{KNIGHT.black(), 4, 4}, {KING.black(), 7, 7}}, true, []Move{0x4436, 0x4456, 0x4432, 0x4452, 0x4425, 0x4423, 0x4463, 0x4465, 0x7766, 0x7776, 0x7767}},
		"Single Knight Obstructed":      {[]TestPiece{{KNIGHT.white(), 4, 4}, {PAWN.white(), 3, 6}, {PAWN.black(), 3, 2}, {KING.white(), 7, 7}}, false, []Move{0x4456, 0x4432, 0x4452, 0x4425, 0x4423, 0x4463, 0x4465, 0x3637, 0x7766, 0x7776, 0x7767}},
		"Discover Check":                {[]TestPiece{{KING.white(), 0, 0}, {PAWN.white(), 1, 1}, {BISHOP.black(), 4, 4}}, false, []Move{0x0001, 0x0010}},
		"Rook Unobstructed":             {[]TestPiece{{ROOK.white(), 3, 3}, {KING.white(), 0, 0}}, false, []Move{0x3330, 0x3331, 0x3332, 0x3334, 0x3335, 0x3336, 0x3337, 0x3303, 0x3313, 0x3323, 0x3343, 0x3353, 0x3363, 0x3373, 0x0001, 0x0010, 0x0011}},
		"Bishop Restricted by Friendly": {[]TestPiece{{BISHOP.black(), 2, 2}, {PAWN.black(), 4, 4}, {KING.black(), 7, 7}}, true, []Move{0x2211, 0x2200, 0x2233, 0x2213, 0x2204, 0x2231, 0x2240, 0x7766, 0x7776, 0x7767, 0x4443}},
		"Absolute Pin":                  {[]TestPiece{{KING.white(), 0, 0}, {ROOK.white(), 1, 1}, {BISHOP.black(), 3, 3}}, false, []Move{0x0001, 0x0010}},
		"King in Double Check":          {[]TestPiece{{KING.white(), 4, 4}, {KNIGHT.black(), 3, 2}, {ROOK.black(), 4, 0}}, false, []Move{0x4433, 0x4434, 0x4435, 0x4454, 0x4455}},
		"Cornered Knight":               {[]TestPiece{{KNIGHT.white(), 0, 0}, {KING.white(), 7, 7}}, false, []Move{0x0012, 0x0021, 0x7766, 0x7767, 0x7776}},
		"Pawn Promotion":                {[]TestPiece{{PAWN.white(), 4, 6}, {KING.white(), 0, 0}}, false, []Move{0x4647, 0x0001, 0x0010, 0x0011}},
		"Queen Mixed": {[]TestPiece{{QUEEN.white(), 3, 3}, {PAWN.white(), 3, 5}, {PAWN.black(), 1, 2}, {KING.white(), 0, 0}}, false, []Move{
			0x3330, 0x3331, 0x3332, 0x3334, 0x3303, 0x3313, 0x3323, 0x3343, 0x3353, 0x3363, 0x3373, 0x3322, 0x3311, 0x3344, 0x3355, 0x3366, 0x3377, 0x3306, 0x3315, 0x3324, 0x3342, 0x3351, 0x3360, 0x0010, 0x0011, 0x3536}},
		"King Cannot Capture Protected Piece": {[]TestPiece{{KING.white(), 4, 4}, {KNIGHT.black(), 5, 5}, {ROOK.black(), 5, 0}}, false, []Move{0x4445, 0x4435, 0x4433}},
	}

	for test := range tests {
		t.Run(test, func(t *testing.T) {
			got := MakeGameGame(tests[test].pieces, [][2]int{})

			got.addPossibleMoves(tests[test].isBlackTurn)

			if !movesMatch(got.possibleMoves, tests[test].expectedMoves) {
				t.Errorf("got\n%v\n but wanted\n%v", got.possibleMoves, tests[test].expectedMoves)
			}
		})
	}
}
