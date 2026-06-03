package gochess

const (
	WHITE uint8 = 0
	BLACK uint8 = 1
)

type Position struct {
	// Bitboards
	pawns   [2]uint64
	knights [2]uint64
	bishops [2]uint64
	rooks   [2]uint64
	queens  [2]uint64
	kings   [2]uint64

	white [2]uint64
	black [2]uint64

	board [2]uint64 // All the pieces

	// Mailbox
	pieces [64]Piece

	// other
	halfmoves uint8
	castle    byte  // Uses first 4 bit for castle right, 1 and 2 are white O-O and O-O-O and 3 and 4 for black
	enPassant uint8 // 0-15 for which square can en passant, and 255 for none
}

func newPosition() *Position {
	pos := &Position{}

	return pos
}
