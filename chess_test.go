package gochess

import "testing"

func assertPiece(t *testing.T, board *ChessBoard, index int, pieceType PieceType, color PieceColor) {
	piece := board.Squares[index].Piece

	if piece == nil || piece.Type != pieceType || piece.Color != color {
		t.Errorf("Expected %v %v at index %v but found %+v", color, pieceType, index, piece)
	}
}

func TestBoardEmptySpace(t *testing.T) {
	board := createBoard()

	for i := 16; i < 48; i++ {
		if board.Squares[i].Piece != nil {
			t.Errorf("Expected empty square at index %v but found %v %v", i, board.Squares[i].Piece.Color, board.Squares[i].Piece.Type)
		}
	}
}

func TestBoardWhitePawns(t *testing.T) {
	board := createBoard()

	for i := 8; i < 16; i++ {
		assertPiece(t, board, i, Pawn, White)
	}
}

func TestBoardWhitePieces(t *testing.T) {
	board := createBoard()

	assertPiece(t, board, 0, Rook, White)
	assertPiece(t, board, 7, Rook, White)
	assertPiece(t, board, 1, Knight, White)
	assertPiece(t, board, 6, Knight, White)
	assertPiece(t, board, 2, Bishop, White)
	assertPiece(t, board, 5, Bishop, White)
	assertPiece(t, board, 3, Queen, White)
	assertPiece(t, board, 4, King, White)
}

func TestBoardBlackPawns(t *testing.T) {
	board := createBoard()

	for i := 48; i < 56; i++ {
		assertPiece(t, board, i, Pawn, Black)
	}
}

func TestBoardBlackPieces(t *testing.T) {
	board := createBoard()

	assertPiece(t, board, 56, Rook, Black)
	assertPiece(t, board, 63, Rook, Black)
	assertPiece(t, board, 57, Knight, Black)
	assertPiece(t, board, 62, Knight, Black)
	assertPiece(t, board, 58, Bishop, Black)
	assertPiece(t, board, 61, Bishop, Black)
	assertPiece(t, board, 59, Queen, Black)
	assertPiece(t, board, 60, King, Black)
}
