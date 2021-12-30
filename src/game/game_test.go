package game_test

import (
	"github.com/martinhrvn/planning-poker/src/game"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGame(t *testing.T) {
	g := game.New()
	assert.Equal(t, g.Started(), false, "game should not start by default")
	assert.Len(t, g.Players(), 0, "game should have empty player slice")
}

func TestStartGame(t *testing.T) {
	g := game.New()
	assert.Equal(t, g.Started(), false, "game should not start by default")
	err := g.Start()
	assert.Nil(t, err, "should be possible to start new game")
	err = g.Start()
	assert.Errorf(t, err, "cannot start already started game")
}

func TestAddPlayers(t *testing.T) {
	tests := []struct {
		name          string
		addPlayers    []game.Player
		resultPlayers []game.Player
	}{
		{name: "add one", addPlayers: []game.Player{{Id: "1", Name: "John"}}, resultPlayers: []game.Player{{Id: "1", Name: "John"}}},
		{
			name:          "add multiple",
			addPlayers:    []game.Player{{Id: "1", Name: "John"}, {Id: "2", Name: "Mark"}},
			resultPlayers: []game.Player{{Id: "1", Name: "John"}, {Id: "2", Name: "Mark"}},
		},
		{
			name:          "add multiple of the same user",
			addPlayers:    []game.Player{{Id: "1", Name: "John"}, {Id: "1", Name: "Mark"}},
			resultPlayers: []game.Player{{Id: "1", Name: "John"}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := game.New()
			g.AddPlayers(test.addPlayers)
			assert.Equal(t, g.Players(), test.resultPlayers, "result length should match")
		})
	}
}
