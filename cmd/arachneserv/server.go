package main

import (
	"io"
	"sync"

	"github.com/pkg/errors"

	pb "github.com/dougfort/arachne/arachne"
	"github.com/dougfort/arachne/internal/game"
)

type arachneServer struct {
	sync.Mutex
	nextID int64
	active map[int64]*game.Game
}

func newServer() *arachneServer {
	var s arachneServer
	s.nextID = 1
	s.active = make(map[int64]*game.Game)
	return &s
}

// Play executes a stream of game commands,
// streaming back the current state of the game
func (a *arachneServer) Play(stream pb.Arachne_PlayServer) error {
	var game *pb.Game

	for {
		playReq, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		switch playReq.GetTestOneof().(type) {
		case *pb.PlayRequest_GameRequest:
			game, err = a.StartGame(playReq.GetGameRequest())
			if err != nil {
				return errors.Wrapf(err, "a.StartGame")
			}
		case *pb.PlayRequest_MoveRequest:
			game, err = a.RequestMove(playReq.GetMoveRequest())
			if err != nil {
				return errors.Wrapf(err, "a.RequestMove")
			}
		case *pb.PlayRequest_DealRequest:
			game, err = a.RequestDeal(playReq.GetDealRequest())
			if err != nil {
				return errors.Wrapf(err, "a.RequestDeal")
			}
		default:
			return errors.Errorf("unknown request")
		}

		if err = stream.Send(game); err != nil {
			return errors.Wrapf(err, "Send()")
		}
	}
}
