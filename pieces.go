package gochess

import "fmt"

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

func getKingThreat(board [8]uint32) [8]uint8 {
	var threat [8]uint8
	return threat
}

func getQueenThreat(board [8]uint32) [8]uint8 {
	var threat [8]uint8
	return threat
}

func getRookThreat(board [8]uint32) [8]uint8 {
	var threat [8]uint8
	return threat
}

func getBishopThreat(board [8]uint32) [8]uint8 {
	var threat [8]uint8
	return threat
}

func getKnightThreat(board [8]uint32) [8]uint8 {
	var threat [8]uint8
	return threat
}

func getPawnThreat(board [8]uint32) [8]uint8 {
	var threat [8]uint8
	return threat
}

func (p Piece) GetThreat(board *Board) ([8]uint8, error) {
	var threat [8]uint8
	if p.Type() == EMPTY {
		return threat, nil
	}
	if p.Type() == PAWN {
		threat = getPawnThreat(board.board)
		return threat, nil
	}
	if p.Type() == KNIGHT {
		threat = getKnightThreat(board.board)
		return threat, nil
	}
	if p.Type() == BISHOP {
		threat = getBishopThreat(board.board)
		return threat, nil
	}
	if p.Type() == ROOK {
		threat = getRookThreat(board.board)
		return threat, nil
	}
	if p.Type() == QUEEN {
		threat = getQueenThreat(board.board)
		return threat, nil
	}
	if p.Type() == KING {
		threat = getKingThreat(board.board)
		return threat, nil
	}
	return threat, fmt.Errorf("Get piece threat failed with invalid piece %v", p.Type())
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

func (p Piece) Strip() Piece {
	return p & 0b00001111
}
