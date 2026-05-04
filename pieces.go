package gochess

type Piece byte

const (
	EMPTY  Piece = 0
	PAWN   Piece = 1
	KNIGHT Piece = 2
	BISHOP Piece = 3
	ROOK   Piece = 4
	QUEEN  Piece = 5
	KING   Piece = 6
)

/*
Get the type of piece stripping extra bits and colour.
*/
func (p Piece) Type() Piece {
	return p & 0b00000111
}

/*
Returns a bool true if the piece is black, false if white or empty.
*/
func (p Piece) IsBlack() bool {
	piece := p.strip()
	piece &= 0b00001000
	return piece > 0
}

/*
Added threats to the board for piece in location. Only if it is their turn
*/
func addThreats(game *Game, x int, y int, isBlackTurn bool) {
	piece, _ := game.GetPiece(x, y)
	if piece.IsBlack() == isBlackTurn {
		return
	}
	switch piece.Type() {
	case EMPTY:
		return
	case PAWN:
		addPawnThreats(game, piece, x, y)
		return
	case KNIGHT:
		addKnightThreats(game, x, y)
		return
	case KING:
		addKingThreats(game, x, y)
		return
	case BISHOP:
		addSlidingPieceThreats(game, x, y, false, true)
		return
	case ROOK:
		addSlidingPieceThreats(game, x, y, true, false)
		return
	case QUEEN:
		addSlidingPieceThreats(game, x, y, true, true)
		return
	}
	game.printThreats()
}

func addSlidingPieceThreats(game *Game, x int, y int, orthogonal bool, diagonal bool) {
	var offsets [][2]int
	if orthogonal {
		offsets = append(offsets, orthogonalOffsets...)
	}
	if diagonal {
		offsets = append(offsets, diagonalOffsets...)
	}
	for i := range offsets {
		targetX := x + offsets[i][0]
		targetY := y + offsets[i][1]
		for targetY <= 7 && targetY >= 0 && targetX <= 7 && targetX >= 0 {
			targetPiece, _ := game.GetPiece(targetX, targetY)
			if targetPiece != EMPTY {
				game.threats[targetY] |= (0b00000001 << targetX)
				break
			}
			game.threats[targetY] |= (0b00000001 << targetX)
			targetX += offsets[i][0]
			targetY += offsets[i][1]
		}
	}
}

func addKingThreats(game *Game, x int, y int) {
	var mask uint16
	if y < 7 {
		mask = (0b00000111 << x) >> 1
		game.threats[y+1] |= uint8(mask)
	}
	if y > 0 {
		mask = (0b00000111 << x) >> 1
		game.threats[y-1] |= uint8(mask)
	}
	mask = (0b00000101 << x) >> 1
	game.threats[y] |= uint8(mask)
}

func addKnightThreats(game *Game, x int, y int) {
	var mask uint16
	if y <= 5 {
		mask = (0b00000101 << x) >> 1
		game.threats[y+2] |= uint8(mask)
	}
	if y >= 2 {
		mask = (0b00000101 << x) >> 1
		game.threats[y-2] |= uint8(mask)
	}
	if y <= 6 {
		mask = (0b00010001 << x) >> 2
		game.threats[y+1] |= uint8(mask)
	}
	if y >= 1 {
		mask = (0b00010001 << x) >> 2
		game.threats[y-1] |= uint8(mask)
	}
}

func addPawnThreats(game *Game, piece Piece, x int, y int) {
	dir := 1
	if piece.IsBlack() {
		dir = -1
	}
	var mask uint16
	mask = (0b00000101 << x) >> 1
	if !(y+dir > 7 || y+dir < 0) {
		game.threats[y+dir] |= uint8(mask)
	}
}

/*
Returns black version of the piece.
*/
func (p Piece) black() Piece {
	return p | 0b00001000
}

/*
Returns white version of the piece.
*/
func (p Piece) white() Piece {
	return p & 0b11110111
}

/*
Returns piece without the 4 extra bits.
*/
func (p Piece) strip() Piece {
	return p & 0b00001111
}
