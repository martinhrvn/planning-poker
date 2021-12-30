package game

import (
	"errors"
	"sort"
)

type Game struct {
	players map[string]Player
	started bool
}

func New() Game {
	return Game{players: map[string]Player{}}
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
	players := make([]Player, 0)
	for _, p := range g.players {
		players = append(players, p)
	}
	sort.Slice(players, func(i, j int) bool {
		return players[i].Id < players[j].Id
	})
	return players
}

func (g *Game) AddPlayers(players []Player) {
	for _, player := range players {
		if _, found := g.players[player.Id]; !found {
			g.players[player.Id] = player
		}
	}
}
