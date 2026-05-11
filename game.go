package gochess

type Game struct {
	position *Position
}

type GameOption func(options *Game)

func BoardLayout(fen FEN) GameOption {
	option := func(options *Game) {
		options.position = fen.getPosition()
	}
	return option
}

func NewGame(opts ...GameOption) *Game {
	game := &Game{
		position: newPosition(),
	}

	for _, opt := range opts {
		opt(game)
	}

	return game
}
