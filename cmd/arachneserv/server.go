package main

import (
	"io"

	"github.com/pkg/errors"

	pb "github.com/dougfort/arachne/arachne"
	"github.com/dougfort/arachne/internal/game"
)

type arachneServer struct{}

func newServer() arachneServer {
	return arachneServer{}
}

// Play executes a stream of game commands,
// streaming back the current state of the game
func (a arachneServer) Play(stream pb.Arachne_PlayServer) error {
	var localGame *game.Game

	for {
		playReq, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		var capture bool

		switch playReq.GetTestOneof().(type) {
		case *pb.PlayRequest_GameRequest:
			if localGame, err = startGame(playReq.GetGameRequest()); err != nil {
				return errors.Wrapf(err, "startGame")
			}
		case *pb.PlayRequest_MoveRequest:
			capture, err = requestMove(localGame, playReq.GetMoveRequest())
			if err != nil {
				return errors.Wrapf(err, "requestMove")
			}
		case *pb.PlayRequest_DealRequest:
			if err = localGame.Deal(); err != nil {
				return errors.Wrapf(err, "Deal")
			}
		default:
			return errors.Errorf("unknown request")
		}

		if err = sendGame(localGame, capture, stream); err != nil {
			return errors.Wrapf(err, "sendGame")
		}
	}
}

func sendGame(localGame *game.Game, capture bool, stream pb.Arachne_PlayServer) error {
	var pbGame pb.Game
	var err error

	pbGame.Seed = localGame.Deck.Seed()
	pbGame.Stack = arachne2pb(localGame.Tableau)
	pbGame.CardsRemaining = int32(localGame.Deck.RemainingCards())
	if capture {
		pbGame.CaptureCount++
	}

	if err = stream.Send(&pbGame); err != nil {
		return errors.Wrapf(err, "Send()")
	}

	return nil
}
