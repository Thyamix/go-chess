package gochess

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type FEN string

const (
	FENDefaultStart FEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)

func (f FEN) getState() (*State, error) {
	fen := string(f)
	if fen == "" {
		return nil, ErrInvalidFen
	}
	parts := strings.Split(fen, " ")
	ranks := strings.Split(parts[0], "/")
	if len(ranks) < 8 || len(ranks) > 8 {
		return nil, ErrFenFenRanks
	}

	state := &State{}

	// Position
	for i, rank := range ranks {
		err := parseRank(rank, i, state)
		if err != nil {
			return nil, err
		}
	}

	state.board = state.white | state.black

	// Active Colour
	switch parts[1] {
	case "b":
		state.IsWhiteTurn = true
	case "w":
		state.IsWhiteTurn = false
	default:
		return nil, ErrFenActiveColor
	}

	// Castle Rights
	castleRight := []rune(parts[2])
	if len(castleRight) > 4 {
		return nil, ErrFenCastleRights
	}
	for _, right := range castleRight {
		switch right {
		case 'k':
			state.castle |= 0b00000001
		case 'q':
			state.castle |= 0b00000010
		case 'K':
			state.castle |= 0b00000100
		case 'Q':
			state.castle |= 0b00001000
		case '-':
			if len(castleRight) > 1 {
				return nil, ErrFenCastleRights
			}
		default:
			return nil, ErrFenCastleRights
		}
	}

	// En Passent
	if parts[3] != "-" {
		enPassentOffset := rune(parts[3][0] - 'a')
		if enPassentOffset > 7 || enPassentOffset < 0 {
			return nil, ErrFenEnPassantTarget
		}

		state.enPassant = state.enPassant | 0x80>>enPassentOffset
	}
	// Move & Halfmove Clock
	halfMoveClock, err := strconv.Atoi(parts[4])
	if err != nil {
		return nil, ErrFenClockInvalid
	}

	moveClock, err := strconv.Atoi(parts[5])
	if err != nil {
		return nil, ErrFenClockInvalid
	}

	if halfMoveClock > (2 * moveClock) {
		return nil, ErrFenHalfMoveToMoveRatio
	}

	if halfMoveClock >= 100 {
		return nil, ErrFenHalfMoveTooHigh
	}

	state.halfMove = halfMoveClock
	state.Move = moveClock

	return state, nil
}

func parseRank(rankString string, rankIndex int, state *State) error {
	pieces := []rune(rankString)

	index := 0 + rankIndex*8
	for _, pieceRune := range pieces {
		isWhite := unicode.IsUpper(pieceRune)
		colourBoard := &state.black
		side := black
		if isWhite {
			colourBoard = &state.white
			side = white
		}
		space := 1
		piece := EMPTY

		shift := 63 - index

		switch unicode.ToLower(pieceRune) {
		case 'p':
			piece = PAWN
			state.pawns[side] = state.pawns[side] | 0x01<<shift
			*colourBoard = *colourBoard | 0x01<<shift
		case 'n':
			piece = KNIGHT
			state.knights[side] = state.knights[side] | 0x01<<shift
			*colourBoard = *colourBoard | 0x01<<shift
		case 'b':
			piece = BISHOP
			state.bishops[side] = state.bishops[side] | 0x01<<shift
			*colourBoard = *colourBoard | 0x01<<shift
		case 'r':
			piece = ROOK
			state.rooks[side] = state.rooks[side] | 0x01<<shift
			*colourBoard = *colourBoard | 0x01<<shift
		case 'q':
			piece = QUEEN
			state.queens[side] = state.queens[side] | 0x01<<shift
			*colourBoard = *colourBoard | 0x01<<shift
		case 'k':
			piece = KING
			state.kings[side] = state.kings[side] | 0x01<<shift
			*colourBoard = *colourBoard | 0x01<<shift
		default:
			space = int(pieceRune - '0')
			if space > 8 {
				return fmt.Errorf("got rune %v of value %v: %w", pieceRune, space, ErrInvalidFen)
			}
		}

		if piece != EMPTY {
			if isWhite {
				piece.toWhite()
			} else {
				piece.toBlack()
			}

			state.Pieces[index] = piece
		}

		index += space

		if index > 64 {
			return fmt.Errorf("too many squares, index %v: %w", index, ErrInvalidFen)
		}
	}
	return nil
}
