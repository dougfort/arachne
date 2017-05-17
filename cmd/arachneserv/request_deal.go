package main

import (
	//	"context"
	context "golang.org/x/net/context"

	"github.com/pkg/errors"

	"github.com/dougfort/arachne/internal/game"

	pb "github.com/dougfort/arachne/arachne"
)

// RequestDeal requests a deal
func (s *arachneServer) RequestDeal(
	ctx context.Context,
	request *pb.DealRequest,
) (*pb.Game, error) {
	var localGame *game.Game
	var pbGame pb.Game
	var ok bool

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if localGame, ok = s.active[request.GetId()]; !ok {
		return nil, errors.Errorf("unknown game id %d", request.GetId())
	}

	if err := localGame.Deal(); err != nil {
		return nil, errors.Wrap(err, "Deal failed")
	}

	pbGame.Seed = localGame.Deck.Seed()
	pbGame.Stack = arachne2pb(localGame.Tableau)
	pbGame.CardsRemaining = int32(localGame.Deck.RemainingCards())

	s.active[pbGame.Id] = localGame
	return &pbGame, nil
}
