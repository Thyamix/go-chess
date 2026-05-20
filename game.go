package gochess

type Game struct {
	position *Position
}

type GameOption func(options *Game) error

func BoardLayout(fen FEN) GameOption {
	position, err := fen.getPosition()
	option := func(options *Game) error {
		if err != nil {
			return err
		}
		options.position = position
		return nil
	}
	return option
}

func NewGame(opts ...GameOption) (*Game, error) {
	game := &Game{
		position: newPosition(),
	}

	for _, opt := range opts {
		err := opt(game)
		if err != nil {
			return nil, err
		}
	}

	return game, nil
}
