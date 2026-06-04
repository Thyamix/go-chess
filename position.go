package gochess

type Color int8
type CastleType int8

const (
	white Color = iota
	black

	shortCastle CastleType = iota
	longCastle
)

// Holds all data held in fen
type State struct {
	// Bitboards
	pawns   [2]uint64
	knights [2]uint64
	bishops [2]uint64
	rooks   [2]uint64
	queens  [2]uint64
	kings   [2]uint64

	white uint64
	black uint64

	board uint64 // All the pieces

	// Mailbox
	Pieces [64]Piece

	// other
	IsWhiteTurn bool
	halfmoves   uint8
	castle      byte  // Uses first 4 bit for castle right, 1 and 2 are white O-O and O-O-O and 3 and 4 for black
	enPassant   uint8 // 0-15 for which square can en passant, and 255 for none

	halfMove int
	Move     int
}

func newState() *State {
	state := &State{}

	return state
}

func (s *State) canCastle(castleType CastleType, color Color) bool {
	if color == white {
		if castleType == shortCastle {
			return s.castle&0b00000001 > 0
		} else {
			return s.castle&0b00000010 > 0
		}
	} else {
		if castleType == shortCastle {
			return s.castle&0b00000100 > 0
		} else {
			return s.castle&0b00001000 > 0
		}
	}
}
