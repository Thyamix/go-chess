package gochess

import "fmt"

func NewBoard() [8]uint32 {
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
	return board
}

/*
Generate a new board with no pieces on it.
*/
func NewEmptyBoard() [8]uint32 {
	board := [8]uint32{}
	return board
}

/*
Gets the piece a x,y, returns EMPTY if not piece is present or if invalid x,y
*/
func (g *Game) GetPiece(x int, y int) (Piece, error) {
	if x > 7 || y > 7 {
		return EMPTY, fmt.Errorf("x: %v and y: %v is not a valid square", x, y)
	}
	row := g.board[y]
	piece := (Piece)(row>>(4*x)) & 0xF
	return piece, nil
}

func (g *Game) isInCheck(isBlack bool) bool {
	var king [2]int
	if isBlack {
		king = g.blackKing
	} else {
		king = g.whiteKing
	}
	if ((g.threats[king[1]] >> king[0]) & 0b00000001) > 0 {
		return true
	}
	return false
}

/*
Gets the threat map based on player turn.
If isBlackTurn == True then it will show all of the black threats and vice versa
*/
func (g *Game) getThreats(isBlackTurn bool) {
	for x := range 8 {
		for y := range 8 {
			addThreats(g, x, y, isBlackTurn)
		}
	}
}

/*
Display the threats in std out to help with debugging.
*/
func (g *Game) printThreats() {
	for y := 7; y >= 0; y-- {
		for x := range 8 {
			bit := (g.threats[y] >> x) & 1
			if bit == 1 {
				if p, _ := g.GetPiece(x, y); p == EMPTY {
					fmt.Print("X ")
				} else {
					fmt.Print("O ")
				}
			} else {
				p, _ := g.GetPiece(x, y)
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
func (g *Game) setPiece(piece Piece, x int, y int) {
	piece = piece.strip()

	//Clear existing piece
	g.board[y] = g.board[y] & ^(0x0000000F << (x * 4))
	//Set new piece
	g.board[y] = g.board[y] | ((0x00000000 | uint32(piece)) << (x * 4))
}
