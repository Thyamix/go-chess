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

func addThreats(board *Board, x int, y int, isBlackTurn bool) {
	piece, _ := board.GetPiece(x, y)
	if piece.IsBlack() != isBlackTurn {
		return
	}
	switch piece.Type() {
	case EMPTY:
		return
	case PAWN:
		addPawnThreats(board, piece, x, y)
		return
	case KNIGHT:
		addKnightThreats(board, piece, x, y)
		return
	}
}

func addKnightThreats(board *Board, piece Piece, x int, y int) {
	var mask uint16
	if y <= 5 {
		mask = (0b00000101 << x) >> 1
		board.threats[y+2] |= uint8(mask)
	}
	if y >= 2 {
		mask = (0b00000101 << x) >> 1
		board.threats[y-2] |= uint8(mask)
	}
	if y <= 6 {
		mask = (0b00010001 << x) >> 2
		board.threats[y+1] |= uint8(mask)
	}
	if y >= 1 {
		mask = (0b00010001 << x) >> 2
		board.threats[y-1] |= uint8(mask)
	}
}

func addPawnThreats(board *Board, piece Piece, x int, y int) {
	dir := 1
	if piece.IsBlack() {
		dir = -1
	}
	var mask uint16
	mask = (0b00000101 << x) >> 1
	board.threats[y+dir] |= uint8(mask)
}

func (p Piece) Type() Piece {
	return p & 0b00000111
}

func (p Piece) Black() Piece {
	return p | 0b00001000
}

func (p Piece) White() Piece {
	return p & 0b11110111
}

func (p Piece) IsBlack() bool {
	piece := p.Strip()
	piece &= 0b00001000
	return piece > 0
}

func (p Piece) Strip() Piece {
	return p & 0b00001111
}
