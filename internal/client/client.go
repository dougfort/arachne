package client

import (
	"context"

	"google.golang.org/grpc"

	"github.com/pkg/errors"

	"github.com/dougfort/gocards"

	pb "github.com/dougfort/arachne/arachne"
	"github.com/dougfort/arachne/internal/game"
)

// LocalGame is the client's representation of the state of the game
type LocalGame struct {
	Tableau        game.Tableau
	CardsRemaining int
	Seed           int64
	CaptureCount   int
}

// Client manages communication wiht the arachne server
type Client interface {

	// NewGame requests a new game with a random seed
	NewGame() (LocalGame, error)

	// ReplayGame requests a new game with a known seed
	ReplayGame(seed int64) (LocalGame, error)

	// Move a card or cards from one stack to another
	Move(move game.MoveType) (LocalGame, error)

	// Deal the next round of cards
	Deal() (LocalGame, error)

	// Close closes the server connection and releasess all resources
	Close() error
}

type clientImpl struct {
	conn     *grpc.ClientConn
	pbClient pb.ArachneClient
	stream   pb.Arachne_PlayClient
}

// New returns an entity that implements the Client interface
func New(address string) (Client, error) {
	var opts []grpc.DialOption
	var c clientImpl
	var err error

	opts = append(opts, grpc.WithInsecure())
	c.conn, err = grpc.Dial(address, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to dial: %s", address)
	}
	c.pbClient = pb.NewArachneClient(c.conn)
	if c.stream, err = c.pbClient.Play(context.Background()); err != nil {
		return nil, errors.Wrap(err, "c.pbClient.Play")
	}

	return &c, nil
}

// NewGame requests a new game with a random seed
func (c *clientImpl) NewGame() (LocalGame, error) {
	var pbGame *pb.Game
	var gameReq pb.GameRequest
	var playGameReq pb.PlayRequest_GameRequest
	var playReq pb.PlayRequest
	var lg LocalGame
	var err error

	gameReq.Gametype = pb.GameRequest_RANDOM
	playGameReq.GameRequest = &gameReq
	playReq.TestOneof = &playGameReq

	if err = c.stream.Send(&playReq); err != nil {
		return LocalGame{}, errors.Wrap(err, "c.stream.Send")
	}

	if pbGame, err = c.stream.Recv(); err != nil {
		return LocalGame{}, errors.Wrap(err, "c.stream.Recv()")
	}

	lg.Seed = pbGame.Seed
	lg.CardsRemaining = int(pbGame.CardsRemaining)

	if lg.Tableau, err = pb2arachne(pbGame); err != nil {
		return LocalGame{}, errors.Wrap(err, "pb2arachne")
	}

	return lg, nil
}

// ReplayGame requests a new game with a known seed
func (c *clientImpl) ReplayGame(seed int64) (LocalGame, error) {
	var pbGame *pb.Game
	var gameReq pb.GameRequest
	var playGameReq pb.PlayRequest_GameRequest
	var playReq pb.PlayRequest
	var lg LocalGame
	var err error

	gameReq.Gametype = pb.GameRequest_REPLAY
	gameReq.Seed = seed
	playGameReq.GameRequest = &gameReq
	playReq.TestOneof = &playGameReq

	if err = c.stream.Send(&playReq); err != nil {
		return LocalGame{}, errors.Wrap(err, "c.stream.Send")
	}

	if pbGame, err = c.stream.Recv(); err != nil {
		return LocalGame{}, errors.Wrap(err, "c.stream.Recv()")
	}

	lg.CardsRemaining = int(pbGame.CardsRemaining)

	if lg.Tableau, err = pb2arachne(pbGame); err != nil {
		return LocalGame{}, errors.Wrap(err, "pb2arachne")
	}

	return lg, nil
}

// Move a card or cards from one stack to another
func (c *clientImpl) Move(move game.MoveType) (LocalGame, error) {
	var pbGame *pb.Game
	var moveReq pb.MoveRequest
	var playMoveReq pb.PlayRequest_MoveRequest
	var playReq pb.PlayRequest
	var lg LocalGame
	var err error

	moveReq.FromCol = int32(move.FromCol)
	moveReq.FromRow = int32(move.FromRow)
	moveReq.ToCol = int32(move.ToCol)
	playMoveReq.MoveRequest = &moveReq
	playReq.TestOneof = &playMoveReq

	if err = c.stream.Send(&playReq); err != nil {
		return LocalGame{}, errors.Wrap(err, "c.stream.Send")
	}

	if pbGame, err = c.stream.Recv(); err != nil {
		return LocalGame{}, errors.Wrap(err, "c.stream.Recv()")
	}

	lg.CardsRemaining = int(pbGame.CardsRemaining)
	lg.Seed = pbGame.Seed
	lg.CaptureCount = int(pbGame.CaptureCount)

	if lg.Tableau, err = pb2arachne(pbGame); err != nil {
		return LocalGame{}, errors.Wrap(err, "pb2arachne")
	}

	return lg, nil
}

// Deal the next round of cards
func (c *clientImpl) Deal() (LocalGame, error) {
	var pbGame *pb.Game
	var dealReq pb.DealRequest
	var playDealReq pb.PlayRequest_DealRequest
	var playReq pb.PlayRequest
	var lg LocalGame
	var err error

	playDealReq.DealRequest = &dealReq
	playReq.TestOneof = &playDealReq

	if err = c.stream.Send(&playReq); err != nil {
		return LocalGame{}, errors.Wrap(err, "c.stream.Send")
	}

	if pbGame, err = c.stream.Recv(); err != nil {
		return LocalGame{}, errors.Wrap(err, "c.stream.Recv()")
	}

	lg.CardsRemaining = int(pbGame.CardsRemaining)
	lg.Seed = pbGame.Seed
	lg.CaptureCount = int(pbGame.CaptureCount)

	if lg.Tableau, err = pb2arachne(pbGame); err != nil {
		return LocalGame{}, errors.Wrap(err, "pb2arachne")
	}

	return lg, nil
}

// Close closes the server connection and releasess all resources
func (c *clientImpl) Close() error {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return errors.Wrap(err, "conn.Close()  failed")
		}
		c.pbClient = nil
		c.conn = nil
		c.stream = nil
		return nil
	}
	return errors.Errorf("already Closed")
}

func pb2arachne(pbGame *pb.Game) (game.Tableau, error) {
	var pbStack []*pb.Stack
	var tableau game.Tableau

	pbStack = pbGame.GetStack()
	if len(pbStack) != game.TableauWidth {
		return game.Tableau{},
			errors.Errorf("expecting %d stacks; found %d",
				game.TableauWidth, len(pbStack))
	}

	for col := 0; col < game.TableauWidth; col++ {
		tableau[col].HiddenCount = int(pbStack[col].HiddenCount)
		tableau[col].Cards = make(gocards.Cards, len(pbStack[col].Cards))
		for row := 0; row < len(pbStack[col].Cards); row++ {
			tableau[col].Cards[row].Suit =
				gocards.Suit(pbStack[col].Cards[row].Suit)
			tableau[col].Cards[row].Rank =
				gocards.Rank(pbStack[col].Cards[row].Rank)
		}
	}

	return tableau, nil
}
