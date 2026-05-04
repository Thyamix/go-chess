package gochess

import "fmt"

// 4 bits start x, 4 bits start y, 4 bits end x, 4 bits end y.
type Move uint16

var (
	orthogonalOffsets = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	diagonalOffsets   = [][2]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}
)

func addSlidingPieceMoves(game *Game, x int, y int, orthogonal bool, diagonal bool, isBlack bool) {
	var offsets [][2]int
	if orthogonal {
		offsets = append(offsets, orthogonalOffsets...)
	}
	if diagonal {
		offsets = append(offsets, diagonalOffsets...)
	}
	for i := range offsets {
		targetX := x + offsets[i][0]
		targetY := y + offsets[i][1]
		for targetY <= 7 && targetY >= 0 && targetX <= 7 && targetX >= 0 {
			targetPiece, _ := game.GetPiece(targetX, targetY)
			if targetPiece != EMPTY && targetPiece.IsBlack() == isBlack {
				break
			} else {
				move, err := createMove(x, y, targetX, targetY)
				if err == nil && game.checkMoveLegality(move, isBlack) {
					game.possibleMoves = append(game.possibleMoves, move)
				}
				targetX += offsets[i][0]
				targetY += offsets[i][1]
			}
		}
	}
}

func addKingMoves(game *Game, x int, y int, isBlack bool) {
	directions := [][2]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	for i := range directions {
		targetX := x + directions[i][0]
		targetY := y + directions[i][1]
		if targetY <= 7 && targetY >= 0 && targetX <= 7 && targetX >= 0 {
			targetPiece, _ := game.GetPiece(targetX, targetY)
			move, err := createMove(x, y, targetX, targetY)
			if targetPiece != EMPTY && targetPiece.IsBlack() == isBlack {
				continue
			}
			if err == nil && game.checkMoveLegality(move, isBlack) {
				game.possibleMoves = append(game.possibleMoves, move)
			}
		}
	}
}

func addKnightMoves(game *Game, x int, y int, isBlack bool) {
	directions := [][2]int{{-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}
	for i := range directions {
		targetX := x + directions[i][0]
		targetY := y + directions[i][1]
		if targetY <= 7 && targetY >= 0 && targetX <= 7 && targetX >= 0 {
			targetPiece, _ := game.GetPiece(targetX, targetY)
			move, err := createMove(x, y, targetX, targetY)
			if targetPiece != EMPTY && targetPiece.IsBlack() == isBlack {
				continue
			}
			if err == nil && game.checkMoveLegality(move, isBlack) {
				game.possibleMoves = append(game.possibleMoves, move)
			}
		}
	}
}

func addPawnMoves(game *Game, x int, y int, isBlack bool) {
	dir := 1
	homerank := 1
	if isBlack {
		dir = -1
		homerank = 6
	}
	if piece, _ := game.GetPiece(x, y+dir); piece == EMPTY {
		if y+dir <= 7 && y+dir >= 0 {
			move, err := createMove(x, y, x, y+dir)
			if err != nil {
				return
			}
			if game.checkMoveLegality(move, isBlack) {
				game.possibleMoves = append(game.possibleMoves, move)
			}
		}
		if piece, _ := game.GetPiece(x, y+dir*2); piece == EMPTY && y == homerank {
			if y+dir <= 7 && y+dir*2 >= 0 {
				move, err := createMove(x, y, x, y+dir*2)
				if err == nil {
					if game.checkMoveLegality(move, isBlack) {
						game.possibleMoves = append(game.possibleMoves, move)
					}
				}
			}
		}
	}
	if piece, _ := game.GetPiece(x+1, y+dir); piece != EMPTY && piece.IsBlack() != isBlack {
		if y+dir <= 7 && y+dir >= 0 {
			move, err := createMove(x, y, x+1, y+dir)
			if err != nil {
				return
			}
			if game.checkMoveLegality(move, isBlack) {
				game.possibleMoves = append(game.possibleMoves, move)
			}
		}
	}
	if piece, _ := game.GetPiece(x-1, y+dir); piece != EMPTY && piece.IsBlack() != isBlack {
		if y+dir <= 7 && y+dir >= 0 {
			move, err := createMove(x, y, x-1, y+dir)
			if err != nil {
				return
			}
			if game.checkMoveLegality(move, isBlack) {
				game.possibleMoves = append(game.possibleMoves, move)
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
Checks all threats to check move legality and to make sure king is not threatened.
*/
func (g *Game) checkMoveLegality(move Move, isBlackTurn bool) bool {
	testGame := *g
	testGame.execMove(move)
	testGame.getThreats(isBlackTurn)

	return !testGame.isInCheck(isBlackTurn)
}

/*
Moves piece by setting EMPTY to current location, and setPiece to new (this replaces any piece currently there)
*/
func (g *Game) execMove(move Move) error {
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

	piece, err := g.GetPiece(x, y)
	if err != nil {
		return err
	}
	g.setPiece(EMPTY, x, y)
	g.setPiece(piece, destX, destY)

	if piece.Type() == KING {
		if piece.IsBlack() {
			g.blackKing = [2]int{destX, destY}
		} else {
			g.whiteKing = [2]int{destX, destY}
		}
	}

	return nil
}

/*
Adds all possible moves to Game for the selected turn. Black for true, White for false.
*/
func (g *Game) addPossibleMoves(isBlackTurn bool) {
	for y := range 8 {
		for x := range 8 {
			piece, _ := g.GetPiece(x, y)
			if piece.IsBlack() == isBlackTurn {
				switch piece.Type() {
				case PAWN:
					addPawnMoves(g, x, y, isBlackTurn)
				case KNIGHT:
					addKnightMoves(g, x, y, isBlackTurn)
				case KING:
					addKingMoves(g, x, y, isBlackTurn)
				case BISHOP:
					addSlidingPieceMoves(g, x, y, false, true, isBlackTurn)
				case ROOK:
					addSlidingPieceMoves(g, x, y, true, false, isBlackTurn)
				case QUEEN:
					addSlidingPieceMoves(g, x, y, true, true, isBlackTurn)
				}
			}

		}
	}
}
