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

func (p Piece) Type() Piece {
	return p & 0b00000111
}

func (p Piece) Black() Piece {
	return p | 0b00001000
}

func (p Piece) White() Piece {
	return p & 0b11110111
}

func (p Piece) Strip() Piece {
	return p & 0b00001111
}
