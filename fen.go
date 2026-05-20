package gochess

import (
	"fmt"
	"strings"
)

type FEN string

const (
	StartingPositionFEN FEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)

func (f FEN) getPosition() (*Position, error) {
	// TODO: Load pos from FEN
	fen := string(f)
	if fen == "" {
		return nil, fmt.Errorf("FEN cannot be empty string")
	}
	parts := strings.Split(fen, " ")
	ranks := strings.Split(parts[0], "/")
	if len(ranks) < 8 {
		return nil, fmt.Errorf("FEN must contain all 8 ranks")
	}

	pos := &Position{}

	for _, rank := range ranks {
		parseRank(rank, pos)
	}

	return pos, nil
}

func parseRank(rank string, pos *Position) {}
