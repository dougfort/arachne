package main

import (
	"github.com/dougfort/arachne/game"
)

func newGame() (gameData, error) {
	var g gameData

	g.remote = game.New()
	return g, nil
}
