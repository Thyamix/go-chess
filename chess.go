package gochess

type Game struct {
	board         [8]uint32
	threats       [8]byte
	possibleMoves []Move
	whiteKing     [2]int
	blackKing     [2]int
	moves         []Move
}

/*
Generate a new Game with board position of board.
*/
func NewGame(board [8]uint32) Game {
	return Game{board: board, whiteKing: [2]int{3, 0}, blackKing: [2]int{3, 7}}
}
