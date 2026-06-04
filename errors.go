package gochess

import (
	"errors"
	"fmt"
)

// Fen Errors
var (
	ErrInvalidFen             = errors.New("invalid fen")
	ErrFenCastleRights        = fmt.Errorf("invalid castle rights: %w", ErrInvalidFen)
	ErrFenFenRanks            = fmt.Errorf("fen must contain all 8 ranks: %w", ErrInvalidFen)
	ErrFenActiveColor         = fmt.Errorf("active color is invalid, only accept 'w' or 'b': %w", ErrInvalidFen)
	ErrFenEnPassantTarget     = fmt.Errorf("invalid en passant target: %w", ErrInvalidFen)
	ErrFenClockInvalid        = fmt.Errorf("half move / move clock invalid: %w", ErrInvalidFen)
	ErrFenHalfMoveTooHigh     = fmt.Errorf("half move clock too high: %w", ErrInvalidFen)
	ErrFenHalfMoveToMoveRatio = fmt.Errorf("half move clock is more then double move clock: %w", ErrInvalidFen)
)
