package gochess

import "testing"

func TestNewGame(t *testing.T) {
	var tests map[string]struct {
		opts     []GameOption
		expected *Game
	}
	for test := range tests {
		t.Run(test, func(t *testing.T) {
			game, err := NewGame(tests[test].opts...)
			if err != nil {
				t.Error(err)
			}
			if game != tests[test].expected {
				t.Errorf("wanted %v, got %v", tests[test].expected, game)
			}
		})
	}
}
