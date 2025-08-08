package gochess

type PieceType uint8
type PieceColor uint8

const (
	Pawn PieceType = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

const (
	White PieceColor = iota
	Black
)

type ChessBoard struct {
	Squares [64]Square
}

type ChessPiece struct {
	Color    PieceColor
	Type     PieceType
	HasMoved bool
}

type Square struct {
	PossibleMoves []Move
	Piece         *ChessPiece
}

type Move struct {
	IsLegal     bool
	Piece       ChessPiece
	IsCastle    bool
	IsEnPassant bool
}

func makeChessPiece(pieceType PieceType, color PieceColor) *ChessPiece {
	return &ChessPiece{
		Type:     pieceType,
		Color:    color,
		HasMoved: false,
	}
}

func makeSquare(piece *ChessPiece) Square {
	return Square{
		PossibleMoves: nil,
		Piece:         piece,
	}
}

func createBoard() *ChessBoard {
	var board ChessBoard

	generateBackRow(&board, White)
	for i := range 8 {
		board.Squares[i+8] = makeSquare(makeChessPiece(Pawn, White))
	}
	for i := range 32 {
		board.Squares[i+16] = makeSquare(nil)
	}
	for i := range 8 {
		board.Squares[i+48] = makeSquare(makeChessPiece(Pawn, Black))
	}
	generateBackRow(&board, Black)

	return &board
}

func generateBackRow(board *ChessBoard, color PieceColor) {
	pos := 0
	if color == Black {
		pos = 56
	}

	board.Squares[pos] = makeSquare(makeChessPiece(Rook, color))
	pos += 1

	board.Squares[pos] = makeSquare(makeChessPiece(Knight, color))
	pos += 1

	board.Squares[pos] = makeSquare(makeChessPiece(Bishop, color))
	pos += 1

	board.Squares[pos] = makeSquare(makeChessPiece(Queen, color))
	pos += 1

	board.Squares[pos] = makeSquare(makeChessPiece(King, color))
	pos += 1

	board.Squares[pos] = makeSquare(makeChessPiece(Bishop, color))
	pos += 1

	board.Squares[pos] = makeSquare(makeChessPiece(Knight, color))
	pos += 1

	board.Squares[pos] = makeSquare(makeChessPiece(Rook, color))
}
