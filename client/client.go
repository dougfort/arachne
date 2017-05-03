package client

import (
	"context"

	"github.com/pkg/errors"

	"google.golang.org/grpc"

	pb "github.com/dougfort/arachne/arachne"

	"github.com/dougfort/arachne/game"
	"github.com/dougfort/gocards"
)

// Client manages communication wiht the arachne server
type Client interface {

	// NewGame requests a new game with a random seed
	NewGame() (game.Tableau, error)

	// ReplayGame requests a new game with a known seed
	ReplayGame(seed int64) (game.Tableau, error)

	// Close closes the server connection and releasess all resources
	Close() error
}

type clientImpl struct {
	conn     *grpc.ClientConn
	pbClient pb.ArachneClient
	gameID   int64
	seed     int64
}

const serverAddr = "127.0.0.1:10000"

// New returns an entity that implements the Client interface
func New() (Client, error) {
	var opts []grpc.DialOption
	var c clientImpl
	var err error

	opts = append(opts, grpc.WithInsecure())
	c.conn, err = grpc.Dial(serverAddr, opts...)
	if err != nil {
		return nil, errors.WithMessage(err, "fail to dial")
	}
	c.pbClient = pb.NewArachneClient(c.conn)

	return &c, nil
}

// NewGame requests a new game with a random seed
func (c *clientImpl) NewGame() (game.Tableau, error) {
	var pbGame *pb.Game
	var err error

	pbGame, err = c.pbClient.StartGame(context.Background(), &pb.GameRequest{})
	if err != nil {
		return game.Tableau{}, errors.Wrap(err, "NewGame()")
	}

	return pb2arachne(pbGame)
}

// ReplayGame requests a new game with a known seed
func (c *clientImpl) ReplayGame(seed int64) (game.Tableau, error) {
	var pbGame *pb.Game
	var err error

	pbGame, err = c.pbClient.StartGame(
		context.Background(),
		&pb.GameRequest{
			Gametype: pb.GameRequest_REPLAY,
			Seed:     seed,
		},
	)
	if err != nil {
		return game.Tableau{}, errors.Wrap(err, "ReplayGame()")
	}

	return pb2arachne(pbGame)
}

// Close closes the server connection and releasess all resources
func (c *clientImpl) Close() error {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return errors.Wrap(err, "conn.Close()  failed")
		}
		c.pbClient = nil
		c.conn = nil
		c.gameID = 0
		c.seed = 0
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
