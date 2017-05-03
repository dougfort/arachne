package main

import (
	//	"context"
	context "golang.org/x/net/context"

	"github.com/pkg/errors"

	"github.com/dougfort/arachne/game"

	pb "github.com/dougfort/arachne/arachne"
)

// RequestMove requests a move of cards
func (s *arachneServer) RequestMove(
	ctx context.Context,
	request *pb.MoveRequest,
) (*pb.Game, error) {
	var move game.MoveType
	var localGame *game.Game
	var pbGame pb.Game
	var ok bool

	move.FromCol = int(request.FromRow)
	move.FromRow = int(request.FromRow)
	move.ToCol = int(request.ToCol)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if localGame, ok = s.active[request.GetId()]; !ok {
		return nil, errors.Errorf("unknown game id %d", request.GetId())
	}

	if err := localGame.Move(move); err != nil {
		return nil, errors.Wrapf(err, "Move %s failed", move)
	}

	pbGame.Seed = localGame.Deck.Seed()
	pbGame.Stack = arachne2pb(localGame.Tableau)
	pbGame.CardsRemaining = int32(localGame.Deck.RemainingCards())

	s.active[pbGame.Id] = localGame
	return &pbGame, nil
}
