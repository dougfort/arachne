package main

import (
	//	"context"
	"fmt"
	"net"

	context "golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/dougfort/arachne/game"

	pb "github.com/dougfort/arachne/arachne"
)

const port = 10000

type arachneServer struct {
	nextID int64
	active map[int64]*game.Game
}

// StartGame starts a new game
func (s *arachneServer) StartGame(
	ctx context.Context,
	request *pb.GameRequest,
) (*pb.Game, error) {
	var localGame *game.Game
	var pbGame pb.Game

	pbGame.Id = s.nextID
	s.nextID++

	switch request.Gametype {
	case pb.GameRequest_RANDOM:
		localGame = game.New()
	case pb.GameRequest_REPLAY:
		localGame = game.Replay(request.Seed)
	default:
		return nil, fmt.Errorf("invalid GameType %d", request.Gametype)
	}

	pbGame.Seed = localGame.Deck.Seed()
	pbGame.Stack = arachne2pb(localGame.Tableau)
	pbGame.CardsRemaining = int32(localGame.Deck.RemainingCards())

	s.active[pbGame.Id] = localGame
	return &pbGame, nil
}

func arachne2pb(tableau game.Tableau) []*pb.Stack {
	stack := make([]*pb.Stack, game.TableauWidth)
	for col := 0; col < game.TableauWidth; col++ {
		stack[col] = new(pb.Stack)
		stack[col].HiddenCount = int32(tableau[col].HiddenCount)
		cardsLen := len(tableau[col].Cards)
		stack[col].Cards = make([]*pb.Card, cardsLen)
		for row := 0; row < cardsLen; row++ {
			localCard := tableau[col].Cards[row]
			stack[col].Cards[row] =
				&pb.Card{
					Suit: int32(localCard.Suit),
					Rank: int32(localCard.Rank),
				}
		}
	}

	return stack
}

func newServer() *arachneServer {
	var s arachneServer
	s.nextID = 1
	s.active = make(map[int64]*game.Game)
	return &s
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterArachneServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
