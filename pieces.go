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

func addBishopMoves(board *Board, x int, y int, isBlack bool) {
	// Top right
	targetX := x + 1
	targetY := y + 1
	finished := false
	for targetY <= 7 && targetY >= 0 && targetX <= 7 && targetX >= 0 && !finished {
		targetPiece, _ := board.GetPiece(targetX, targetY)
		if targetPiece != EMPTY || targetPiece.IsBlack() == isBlack {
			finished = true
		} else {
			move, err := createMove(x, y, targetX, targetY)
			if err == nil {
				board.possibleMoves = append(board.possibleMoves, move)
			}
			targetX += 1
			targetY += 1
		}
	}
	// Bottom right
	targetX = x + 1
	targetY = y - 1
	finished = false
	for targetY <= 7 && targetY >= 0 && targetX <= 7 && targetX >= 0 && !finished {
		targetPiece, _ := board.GetPiece(targetX, targetY)
		if targetPiece != EMPTY || targetPiece.IsBlack() == isBlack {
			finished = true
		} else {
			move, err := createMove(x, y, targetX, targetY)
			if err == nil {
				board.possibleMoves = append(board.possibleMoves, move)
			}
			targetX += 1
			targetY -= 1
		}
	}
	// Top left
	targetX = x - 1
	targetY = y + 1
	finished = false
	for targetY <= 7 && targetY >= 0 && targetX <= 7 && targetX >= 0 && !finished {
		targetPiece, _ := board.GetPiece(targetX, targetY)
		if targetPiece != EMPTY || targetPiece.IsBlack() == isBlack {
			finished = true
		} else {
			move, err := createMove(x, y, targetX, targetY)
			if err == nil {
				board.possibleMoves = append(board.possibleMoves, move)
			}
			targetX -= 1
			targetY += 1
		}
	}
	// Botton left
	targetX = x - 1
	targetY = y - 1
	finished = false
	for targetY <= 7 && targetY >= 0 && targetX <= 7 && targetX >= 0 && !finished {
		targetPiece, _ := board.GetPiece(targetX, targetY)
		if targetPiece != EMPTY || targetPiece.IsBlack() == isBlack {
			finished = true
		} else {
			move, err := createMove(x, y, targetX, targetY)
			if err == nil {
				board.possibleMoves = append(board.possibleMoves, move)
			}
			targetX -= 1
			targetY -= 1
		}
	}
}

func addKingMoves(board *Board, x int, y int, isBlack bool) {
	// Right
	move, err := createMove(x, y, x+1, y)
	if err == nil {
		if piece, _ := board.GetPiece(x+1, y); piece.IsBlack() != isBlack || piece.Type() == EMPTY {
			if board.CheckMoveLegality(move, isBlack) {
				board.possibleMoves = append(board.possibleMoves, move)
			}
		}
	}
	// Top Right
	move, err = createMove(x, y, x+1, y+1)
	if err == nil {
		if piece, _ := board.GetPiece(x+1, y+1); piece.IsBlack() != isBlack || piece.Type() == EMPTY {
			if board.CheckMoveLegality(move, isBlack) {
				board.possibleMoves = append(board.possibleMoves, move)
			}
		}
	}
	// Bottom Right
	move, err = createMove(x, y, x+1, y-1)
	if err == nil {
		if piece, _ := board.GetPiece(x+1, y-1); piece.IsBlack() != isBlack || piece.Type() == EMPTY {
			if board.CheckMoveLegality(move, isBlack) {
				board.possibleMoves = append(board.possibleMoves, move)
			}
		}
	}
	// Left
	move, err = createMove(x, y, x-1, y)
	if err == nil {
		if piece, _ := board.GetPiece(x-1, y); piece.IsBlack() != isBlack || piece.Type() == EMPTY {
			if board.CheckMoveLegality(move, isBlack) {
				board.possibleMoves = append(board.possibleMoves, move)
			}
		}
	} // Top Left
	move, err = createMove(x, y, x-1, y+1)
	if err == nil {
		if piece, _ := board.GetPiece(x-1, y+1); piece.IsBlack() != isBlack || piece.Type() == EMPTY {
			if board.CheckMoveLegality(move, isBlack) {
				board.possibleMoves = append(board.possibleMoves, move)
			}
		}
	}
	// Bottom Left
	move, err = createMove(x, y, x-1, y-1)
	if err == nil {
		if piece, _ := board.GetPiece(x-1, y-1); piece.IsBlack() != isBlack || piece.Type() == EMPTY {
			if board.CheckMoveLegality(move, isBlack) {
				board.possibleMoves = append(board.possibleMoves, move)
			}
		}
	}
	// Top
	move, err = createMove(x, y, x, y+1)
	if err == nil {
		if piece, _ := board.GetPiece(x, y+1); piece.IsBlack() != isBlack || piece.Type() == EMPTY {
			if board.CheckMoveLegality(move, isBlack) {
				board.possibleMoves = append(board.possibleMoves, move)
			}
		}
	}
	// Bottom
	move, err = createMove(x, y, x, y-1)
	if err == nil {
		if piece, _ := board.GetPiece(x, y-1); piece.IsBlack() != isBlack || piece.Type() == EMPTY {
			if board.CheckMoveLegality(move, isBlack) {
				board.possibleMoves = append(board.possibleMoves, move)
			}
		}
	}
}

func addKnightMoves(board *Board, x int, y int, isBlack bool) {
	// Right
	move, err := createMove(x, y, x+2, y+1)
	if err == nil {
		if piece, _ := board.GetPiece(x+2, y+1); piece.IsBlack() != isBlack || piece.Type() == EMPTY {
			if board.CheckMoveLegality(move, isBlack) {
				board.possibleMoves = append(board.possibleMoves, move)
			}
		}
	}
	move, err = createMove(x, y, x+2, y-1)
	if err == nil {
		if piece, _ := board.GetPiece(x+2, y-1); piece.IsBlack() != isBlack || piece.Type() == EMPTY {
			if board.CheckMoveLegality(move, isBlack) {
				board.possibleMoves = append(board.possibleMoves, move)
			}
		}
	}
	// Left
	move, err = createMove(x, y, x-2, y+1)
	if err == nil {
		if piece, _ := board.GetPiece(x-2, y+1); piece.IsBlack() != isBlack || piece.Type() == EMPTY {
			if board.CheckMoveLegality(move, isBlack) {
				board.possibleMoves = append(board.possibleMoves, move)
			}
		}
	}
	move, err = createMove(x, y, x-2, y-1)
	if err == nil {
		if piece, _ := board.GetPiece(x-2, y-1); piece.IsBlack() != isBlack || piece.Type() == EMPTY {
			if board.CheckMoveLegality(move, isBlack) {
				board.possibleMoves = append(board.possibleMoves, move)
			}
		}
	}
	// Top
	move, err = createMove(x, y, x+1, y+2)
	if err == nil {
		if piece, _ := board.GetPiece(x+1, y+2); piece.IsBlack() != isBlack || piece.Type() == EMPTY {
			if board.CheckMoveLegality(move, isBlack) {
				board.possibleMoves = append(board.possibleMoves, move)
			}
		}
	}
	move, err = createMove(x, y, x-1, y+2)
	if err == nil {
		if piece, _ := board.GetPiece(x-1, y+2); piece.IsBlack() != isBlack || piece.Type() == EMPTY {
			if board.CheckMoveLegality(move, isBlack) {
				board.possibleMoves = append(board.possibleMoves, move)
			}
		}
	}
	// Bottom
	move, err = createMove(x, y, x+1, y-2)
	if err == nil {
		if piece, _ := board.GetPiece(x+1, y-2); piece.IsBlack() != isBlack || piece.Type() == EMPTY {
			if board.CheckMoveLegality(move, isBlack) {
				if board.CheckMoveLegality(move, isBlack) {
					board.possibleMoves = append(board.possibleMoves, move)
				}
			}
		}
	}
	move, err = createMove(x, y, x-1, y-2)
	if err == nil {
		if piece, _ := board.GetPiece(x-1, y-2); piece.IsBlack() != isBlack || piece.Type() == EMPTY {
			if board.CheckMoveLegality(move, isBlack) {
				board.possibleMoves = append(board.possibleMoves, move)
			}
		}
	}

}

func addPawnMoves(board *Board, x int, y int, isBlack bool) {
	dir := 1
	homerank := 1
	if isBlack {
		dir = -1
		homerank = 6
	}
	if piece, _ := board.GetPiece(x, y+dir); piece == EMPTY {
		if y+dir <= 7 && y+dir >= 0 {
			move, err := createMove(x, y, x, y+dir)
			if err != nil {
				return
			}
			if board.CheckMoveLegality(move, isBlack) {
				board.possibleMoves = append(board.possibleMoves, move)
			}
		}
		if piece, _ := board.GetPiece(x, y+dir*2); piece == EMPTY && y == homerank {
			if y+dir <= 7 && y+dir*2 >= 0 {
				move, err := createMove(x, y, x, y+dir*2)
				if err == nil {
					if board.CheckMoveLegality(move, isBlack) {
						board.possibleMoves = append(board.possibleMoves, move)
					}
				}
			}
		}
	}
	if piece, _ := board.GetPiece(x+1, y+dir); piece != EMPTY && piece.IsBlack() != isBlack {
		if y+dir <= 7 && y+dir >= 0 {
			move, err := createMove(x, y, x+1, y+dir)
			if err != nil {
				return
			}
			if board.CheckMoveLegality(move, isBlack) {
				board.possibleMoves = append(board.possibleMoves, move)
			}
		}
	}
	if piece, _ := board.GetPiece(x-1, y+dir); piece != EMPTY && piece.IsBlack() != isBlack {
		if y+dir <= 7 && y+dir >= 0 {
			move, err := createMove(x, y, x-1, y+dir)
			if err != nil {
				return
			}
			if board.CheckMoveLegality(move, isBlack) {
				board.possibleMoves = append(board.possibleMoves, move)
			}
		}
	}
}

func createMove(pieceX int, pieceY int, targetX int, targetY int) (Move, error) {
	var move Move
	if pieceX > 7 || pieceX < 0 || pieceY > 7 || pieceY < 0 {
		return move, fmt.Errorf("Piece not in play area")
	}
	if targetX > 7 || targetX < 0 || targetY > 7 || targetY < 0 {
		return move, fmt.Errorf("Target not in play area")
	}
	move |= Move(pieceX << 12)
	move |= Move(pieceY << 8)
	move |= Move(targetX << 4)
	move |= Move(targetY)
	return move, nil
}

/*
Added threats to the board for piece in location. Only if it is their turn
*/
func addThreats(board *Board, x int, y int, isBlackTurn bool) {
	piece, _ := board.GetPiece(x, y)
	if piece.IsBlack() == isBlackTurn {
		return
	}
	switch piece.Type() {
	case EMPTY:
		return
	case PAWN:
		addPawnThreats(board, piece, x, y)
		return
	case KNIGHT:
		addKnightThreats(board, x, y)
		return
	case KING:
		addKingThreats(board, x, y)
		return
	case BISHOP:
		addBishopThreats(board, x, y)
		return
	case ROOK:
		addRookThreats(board, x, y)
		return
	case QUEEN:
		addQueenThreats(board, x, y)
		return
	}
	board.printThreats()
}

func addQueenThreats(board *Board, x int, y int) {
	addBishopThreats(board, x, y)
	addRookThreats(board, x, y)
}

func addRookThreats(board *Board, x int, y int) {
	// Top
	targetX := x
	targetY := y + 1
	finished := false
	for targetY <= 7 && targetY >= 0 && targetX <= 7 && targetX >= 0 && !finished {
		targetPiece, _ := board.GetPiece(targetX, targetY)
		if targetPiece != EMPTY {
			finished = true
		}
		board.threats[targetY] |= (0b00000001 << targetX)
		targetY += 1
	}
	// Bottom
	targetX = x
	targetY = y - 1
	finished = false
	for targetY <= 7 && targetY >= 0 && targetX <= 7 && targetX >= 0 && !finished {
		targetPiece, _ := board.GetPiece(targetX, targetY)
		if targetPiece != EMPTY {
			finished = true
		}
		board.threats[targetY] |= (0b00000001 << targetX)
		targetY -= 1
	}
	// Right
	targetX = x + 1
	targetY = y
	finished = false
	for targetY <= 7 && targetY >= 0 && targetX <= 7 && targetX >= 0 && !finished {
		targetPiece, _ := board.GetPiece(targetX, targetY)
		if targetPiece != EMPTY {
			finished = true
		}
		board.threats[targetY] |= (0b00000001 << targetX)
		targetX += 1
	}
	// Left
	targetX = x - 1
	targetY = y
	finished = false
	for targetY <= 7 && targetY >= 0 && targetX <= 7 && targetX >= 0 && !finished {
		targetPiece, _ := board.GetPiece(targetX, targetY)
		if targetPiece != EMPTY {
			finished = true
		}
		board.threats[targetY] |= (0b00000001 << targetX)
		targetX -= 1
	}
}

func addBishopThreats(board *Board, x int, y int) {
	// Top right
	targetX := x + 1
	targetY := y + 1
	finished := false
	for targetY <= 7 && targetY >= 0 && targetX <= 7 && targetX >= 0 && !finished {
		targetPiece, _ := board.GetPiece(targetX, targetX)
		if targetPiece != EMPTY {
			finished = true
		}
		board.threats[targetY] |= (0b00000001 << targetX)
		targetX += 1
		targetY += 1
	}
	// Bottom right
	targetX = x + 1
	targetY = y - 1
	finished = false
	for targetY <= 7 && targetY >= 0 && targetX <= 7 && targetX >= 0 && !finished {
		targetPiece, _ := board.GetPiece(targetX, targetY)
		if targetPiece != EMPTY {
			finished = true
		}
		board.threats[targetY] |= (0b00000001 << targetX)
		targetX += 1
		targetY -= 1
	}
	// Top left
	targetX = x - 1
	targetY = y + 1
	finished = false
	for targetY <= 7 && targetY >= 0 && targetX <= 7 && targetX >= 0 && !finished {
		targetPiece, _ := board.GetPiece(targetX, targetY)
		if targetPiece != EMPTY {
			finished = true
		}
		board.threats[targetY] |= (0b00000001 << targetX)
		targetX -= 1
		targetY += 1
	}
	// Botton left
	targetX = x - 1
	targetY = y - 1
	finished = false
	for targetY <= 7 && targetY >= 0 && targetX <= 7 && targetX >= 0 && !finished {
		targetPiece, _ := board.GetPiece(targetX, targetY)
		if targetPiece != EMPTY {
			finished = true
		}
		board.threats[targetY] |= (0b00000001 << targetX)
		targetX -= 1
		targetY -= 1
	}
}

func addKingThreats(board *Board, x int, y int) {
	var mask uint16
	if y < 7 {
		mask = (0b00000111 << x) >> 1
		board.threats[y+1] |= uint8(mask)
	}
	if y > 0 {
		mask = (0b00000111 << x) >> 1
		board.threats[y-1] |= uint8(mask)
	}
	mask = (0b00000101 << x) >> 1
	board.threats[y] |= uint8(mask)
}

func addKnightThreats(board *Board, x int, y int) {
	var mask uint16
	if y <= 5 {
		mask = (0b00000101 << x) >> 1
		board.threats[y+2] |= uint8(mask)
	}
	if y >= 2 {
		mask = (0b00000101 << x) >> 1
		board.threats[y-2] |= uint8(mask)
	}
	if y <= 6 {
		mask = (0b00010001 << x) >> 2
		board.threats[y+1] |= uint8(mask)
	}
	if y >= 1 {
		mask = (0b00010001 << x) >> 2
		board.threats[y-1] |= uint8(mask)
	}
}

func addPawnThreats(board *Board, piece Piece, x int, y int) {
	dir := 1
	if piece.IsBlack() {
		dir = -1
	}
	var mask uint16
	mask = (0b00000101 << x) >> 1
	if !(y+dir > 7 || y+dir < 0) {
		board.threats[y+dir] |= uint8(mask)
	}
}

/*
Get the type of piece stripping extra bits and colour.
*/
func (p Piece) Type() Piece {
	return p & 0b00000111
}

/*
Returns black version of the piece.
*/
func (p Piece) Black() Piece {
	return p | 0b00001000
}

/*
Returns white version of the piece.
*/
func (p Piece) White() Piece {
	return p & 0b11110111
}

/*
Returns a bool true if the piece is black, false if white or empty.
*/
func (p Piece) IsBlack() bool {
	piece := p.Strip()
	piece &= 0b00001000
	return piece > 0
}

/*
Returns piece without the 4 extra bits.
*/
func (p Piece) Strip() Piece {
	return p & 0b00001111
}
