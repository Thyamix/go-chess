package gochess

import "fmt"

type Board [8]uint32

func NewBoard() Board {
	return Board{
		0b01000010001101100101001100100100,
		0b00010001000100010001000100010001,
		0b00000000000000000000000000000000,
		0b00000000000000000000000000000000,
		0b00000000000000000000000000000000,
		0b00000000000000000000000000000000,
		0b10011001100110011001100110011001,
		0b11001010101111101101101110101100,
	}
}

func (b *Board) Get(x int, y int) (*Piece, error) {
	if x > 7 || y > 7 {
		return nil, fmt.Errorf("x: %v and y: %v is not a valid square", x, y)
	}
	row := b[y]
	piece := (Piece)(row>>(4*x)) & 0xF
	return &piece, nil
}
