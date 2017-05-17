package main

import (
	//	"context"
	context "golang.org/x/net/context"

	"github.com/pkg/errors"

	"github.com/dougfort/arachne/internal/game"

	pb "github.com/dougfort/arachne/arachne"
)

// StartGame starts a new game
func (s *arachneServer) StartGame(
	ctx context.Context,
	request *pb.GameRequest,
) (*pb.Game, error) {
	var localGame *game.Game
	var pbGame pb.Game

	switch request.Gametype {
	case pb.GameRequest_RANDOM:
		localGame = game.New()
	case pb.GameRequest_REPLAY:
		localGame = game.Replay(request.Seed)
	default:
		return nil, errors.Errorf("invalid GameType %d", request.Gametype)
	}

	pbGame.Seed = localGame.Deck.Seed()
	pbGame.Stack = arachne2pb(localGame.Tableau)
	pbGame.CardsRemaining = int32(localGame.Deck.RemainingCards())

	s.mutex.Lock()
	defer s.mutex.Unlock()

	pbGame.Id = s.nextID
	s.nextID++
	s.active[pbGame.Id] = localGame
	return &pbGame, nil
}
