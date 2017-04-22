package main

import (
	"github.com/pkg/errors"

	gamelib "github.com/dougfort/arachne/game"
)

func processMove(g gameData, m gamelib.MoveType) error {
	var err error

	if err = g.remote.Move(m); err != nil {
		return errors.Wrapf(err, "invalid move %s", m)
	}

	return nil
}