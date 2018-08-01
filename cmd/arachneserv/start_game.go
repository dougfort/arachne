package main

import (
	"github.com/pkg/errors"

	"github.com/dougfort/arachne/internal/game"

	pb "github.com/dougfort/arachne/arachne"
)

// StartGame starts a new game
func startGame(request *pb.GameRequest) (*game.Game, error) {
	var localGame *game.Game

	switch request.Gametype {
	case pb.GameRequest_RANDOM:
		localGame = game.New()
	case pb.GameRequest_REPLAY:
		localGame = game.Replay(request.Seed)
	default:
		return nil, errors.Errorf("invalid GameType %d", request.Gametype)
	}

	return localGame, nil
}
