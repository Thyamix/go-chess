package gochess

type FEN string

const (
	StartingPositionFEN FEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
)

func (f FEN) getPosition() *Position {
	// TODO: Load pos from FEN
	return &Position{}
}
