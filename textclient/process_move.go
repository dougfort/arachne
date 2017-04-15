package main

import (
	"github.com/pkg/errors"

	gamelib "github.com/dougfort/arachne/game"
)

func processMove(g gameData, move gamelib.MoveType) error {
	var err error

	if err = g.remote.Tableau.ValidateMove(move); err != nil {
		return errors.Wrapf(err, "invalid move %s", move)
	}

	return nil
}
