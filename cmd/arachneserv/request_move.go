package main

import (
	"github.com/pkg/errors"

	"github.com/dougfort/arachne/internal/game"

	pb "github.com/dougfort/arachne/arachne"
)

// RequestMove requests a move of cards
func requestMove(localGame *game.Game, request *pb.MoveRequest) (bool, error) {
	var move game.MoveType
	var capture bool
	var err error

	move.FromCol = int(request.FromCol)
	move.FromRow = int(request.FromRow)
	move.ToCol = int(request.ToCol)

	if capture, err = localGame.Move(move); err != nil {
		return false, errors.Wrapf(err, "Move %s failed", move)
	}

	return capture, nil
}
