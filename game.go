package gochess

type Game struct {
	State *State
}

type GameOption func(options *Game) error

func BoardLayout(fen FEN) GameOption {
	state, err := fen.getState()
	option := func(options *Game) error {
		if err != nil {
			return err
		}
		options.State = state
		return nil
	}
	return option
}

func NewGame(opts ...GameOption) (*Game, error) {
	game := &Game{
		State: newState(),
	}

	for _, opt := range opts {
		err := opt(game)
		if err != nil {
			return nil, err
		}
	}

	return game, nil
}
