package gochess

import "fmt"

// 4 bits start x, 4 bits start y, 4 bits end x, 4 bits end y.
type Move uint16

type Board struct {
	board         [8]uint32
	threats       [8]byte
	possibleMoves []Move
	whiteKing     [2]int
	blackKing     [2]int
}

/*
Generate a new board with all the pieces in default position.
*/
func NewBoard() Board {
	// NOTE: Board is mirrored, right to left, and white is at top.
	board := [8]uint32{
		0b01000010001101010110001100100100,
		0b00010001000100010001000100010001,
		0b00000000000000000000000000000000,
		0b00000000000000000000000000000000,
		0b00000000000000000000000000000000,
		0b00000000000000000000000000000000,
		0b10011001100110011001100110011001,
		0b11001010101111011110101110101100,
	}
	return Board{board: board, whiteKing: [2]int{3, 0}, blackKing: [2]int{3, 7}}
}

/*
Checks all threats to check move legality and to make sure king is not threatened.
*/
func (b *Board) CheckMoveLegality(move Move, isBlackTurn bool) bool {
	testBoard := *b
	testBoard.ExecMove(move)
	testBoard.GetThreats(isBlackTurn)

	return !testBoard.IsInCheck(isBlackTurn)
}

func (b *Board) IsInCheck(isBlack bool) bool {
	var king [2]int
	if isBlack {
		king = b.blackKing
	} else {
		king = b.whiteKing
	}
	if ((b.threats[king[1]] >> king[0]) & 0b00000001) > 0 {
		return true
	}
	return false
}

/*
Moves piece by setting EMPTY to current location, and setPiece to new (this replaces any piece currently there)
*/
func (b *Board) ExecMove(move Move) error {
	var x int
	var y int
	var destX int
	var destY int
	destY = int(move & 0b0000000000001111)
	move >>= 4
	destX = int(move & 0b0000000000001111)
	move >>= 4
	y = int(move & 0b0000000000001111)
	move >>= 4
	x = int(move & 0b0000000000001111)

	piece, err := b.GetPiece(x, y)
	if err != nil {
		return err
	}
	b.SetPiece(EMPTY, x, y)
	b.SetPiece(piece, destX, destY)

	if piece.Type() == KING {
		if piece.IsBlack() {
			b.blackKing = [2]int{destX, destY}
		} else {
			b.whiteKing = [2]int{destX, destY}
		}
	}

	return nil
}

/*
Adds all possible moves to the Board for the selected turn. Black for true, White for false.
*/
func (b *Board) AddPossibleMoves(isBlackTurn bool) {
	for y := range 8 {
		for x := range 8 {
			piece, _ := b.GetPiece(x, y)
			if piece.IsBlack() == isBlackTurn {
				switch piece.Type() {
				case PAWN:
					addPawnMoves(b, x, y, isBlackTurn)
				case KNIGHT:
					addKnightMoves(b, x, y, isBlackTurn)
				case KING:
					addKingMoves(b, x, y, isBlackTurn)
				case BISHOP:
					addBishopMoves(b, x, y, isBlackTurn)
				}
			}

		}
	}
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
If isBlackTurn == True then it will show all of the black threats and vice versa
*/
func (b *Board) GetThreats(isBlackTurn bool) {
	for x := range 8 {
		for y := range 8 {
			addThreats(b, x, y, isBlackTurn)
		}
	}
}

/*
Display the threats in std out to help with debugging.
*/
func (b *Board) printThreats() {
	for y := 7; y >= 0; y-- {
		for x := range 8 {
			bit := (b.threats[y] >> x) & 1
			if bit == 1 {
				if p, _ := b.GetPiece(x, y); p == EMPTY {
					fmt.Print("X ")
				} else {
					fmt.Print("O ")
				}
			} else {
				p, _ := b.GetPiece(x, y)
				if p != EMPTY {
					fmt.Print("P ")
				} else {
					fmt.Print(". ")
				}
			}
		}
		fmt.Println()
	}
	fmt.Println("----------------")
}

/*
Sets a Piece on the board replacing any other Piece already at that location
*/
func (b *Board) SetPiece(piece Piece, x int, y int) {
	piece = piece.Strip()

	//Clear existing piece
	b.board[y] = b.board[y] & ^(0x0000000F << (x * 4))
	//Set new piece
	b.board[y] = b.board[y] | ((0x00000000 | uint32(piece)) << (x * 4))
}
