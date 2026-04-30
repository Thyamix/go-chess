package gochess

import "fmt"

type Move uint8

type Board struct {
	board   [8]uint32
	threats [8]uint8
	moves   []Move
}

func NewBoard() Board {
	board := [8]uint32{
		0b01000010001101100101001100100100,
		0b00010001000100010001000100010001,
		0b00000000000000000000000000000000,
		0b00000000000000000000000000000000,
		0b00000000000000000000000000000000,
		0b00000000000000000000000000000000,
		0b10011001100110011001100110011001,
		0b11001010101111101101101110101100,
	}
	return Board{board: board}
}

/*
Gets the piece a x,y, returns EMPTY if not piece is present or if invalid x,y
*/
func (b *Board) GetPiece(x int, y int) (Piece, error) {
	if x > 7 || y > 7 {
		return EMPTY, fmt.Errorf("x: %v and y: %v is not a valid square", x, y)
	}
	row := b.board[y]
	piece := (Piece)(row>>(4*x)) & 0xF
	return piece, nil
}

/*
Gets the threat map based on player turn.
If isBlackTurn == True then it will show all of the white threats and vice versa
*/
func (b *Board) GetThreats(isBlackTurn bool) {
	for x := range 8 {
		for y := range 8 {
			addThreats(b, x, y, isBlackTurn)
		}
	}
}

/*
Sets a Piece on the board replacing any other Piece already at that location
*/
func (b *Board) SetPiece(piece Piece, x int, y int) {
	piece = piece.Strip()

	//Clear existing piece
	b.board[y] = b.board[y] & ^(0x0000000F << x * 4)
	//Set new piece
	b.board[y] = b.board[y] | ((0x00000000 | uint32(piece)) << (x * 4))
}
