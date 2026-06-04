package gochess

type Piece uint8

const (
	EMPTY  Piece = 0
	PAWN   Piece = 1
	KNIGHT Piece = 2
	BISHOP Piece = 3
	ROOK   Piece = 4
	QUEEN  Piece = 5
	KING   Piece = 6
)

func (p *Piece) toWhite() {
	*p = *p & 0x07
}

func (p *Piece) toBlack() {
	*p = *p | 0x08
}

func (p *Piece) IsWhite() bool {
	return *p <= 8
}

func (p *Piece) IsBlack() bool {
	return *p > 8
}

func (p *Piece) Type() Piece {
	return *p & 0x07
}
