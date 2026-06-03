package gochess

import (
	"fmt"
	"strings"
	"unicode"
)

type FEN string

const (
	StartingPositionFEN FEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)

func (f FEN) getPosition() (*Position, error) {
	fen := string(f)
	if fen == "" {
		return nil, fmt.Errorf("FEN cannot be empty string")
	}
	parts := strings.Split(fen, " ")
	ranks := strings.Split(parts[0], "/")
	if len(ranks) < 8 || len(ranks) > 8 {
		return nil, fmt.Errorf("FEN must exactly all 8 ranks")
	}

	pos := &Position{}

	for i, rank := range ranks {
		err := parseRank(rank, i, pos)
		if err != nil {
			return nil, err
		}
	}

	return pos, nil
}

func parseRank(rankString string, rankIndex int, pos *Position) error {
	pieces := []rune(rankString)

	index := 0 + rankIndex*8
	for _, pieceRune := range pieces {
		isWhite := unicode.IsUpper(pieceRune)
		side := BLACK
		if isWhite {
			side = WHITE
		}
		space := 1
		piece := EMPTY

		shift := 63 - index

		switch unicode.ToLower(pieceRune) {
		case 'p':
			piece = PAWN
			pos.pawns[side] = pos.pawns[side] | 0x01<<shift
			pos.board[side] = pos.board[side] | 0x01<<shift
		case 'n':
			piece = KNIGHT
			pos.knights[side] = pos.knights[side] | 0x01<<shift
			pos.board[side] = pos.board[side] | 0x01<<shift
		case 'b':
			piece = BISHOP
			pos.bishops[side] = pos.bishops[side] | 0x01<<shift
			pos.board[side] = pos.board[side] | 0x01<<shift
		case 'r':
			piece = ROOK
			pos.rooks[side] = pos.rooks[side] | 0x01<<shift
			pos.board[side] = pos.board[side] | 0x01<<shift
		case 'q':
			piece = QUEEN
			pos.queens[side] = pos.queens[side] | 0x01<<shift
			pos.board[side] = pos.board[side] | 0x01<<shift
		case 'k':
			piece = KING
			pos.kings[side] = pos.kings[side] | 0x01<<shift
			pos.board[side] = pos.board[side] | 0x01<<shift
		default:
			space = int(pieceRune - '0')
			if space > 8 {
				return fmt.Errorf("Invalid, cannot have space greater then 8, got rune %v of value %v", pieceRune, space)
			}
		}

		if piece != EMPTY {
			if isWhite {
				piece.toWhite()
			} else {
				piece.toBlack()
			}

			pos.pieces[index] = piece
		}

		index += space

		if index > 64 {
			return fmt.Errorf("Invalid FEN, too many squares, index: %v", index)
		}
	}
	return nil
}
