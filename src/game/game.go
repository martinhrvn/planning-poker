package game

import "errors"

type Game struct {
	players map[string]Player
	started bool
}

func New() Game {
	return Game{}
}

func (g Game) Started() bool {
	return g.started
}

func (g *Game) Start() error {
	if g.started {
		return errors.New("cannot start already started game")
	}
	g.started = true
	return nil
}

func (g Game) Players() []Player {
	players := make([]Player, len(g.players))
	return players
}

func (g *Game) AddPlayers(players []Player) {
	for _, player := range players {
		g.players[player.Id] = player
	}
}
