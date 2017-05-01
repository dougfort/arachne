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
	pbGame.Stack = make([]*pb.Stack, game.TableauWidth)
	for i := 0; i < game.TableauWidth; i++ {
		pbGame.Stack[i] = new(pb.Stack)
		pbGame.Stack[i].HiddenCount = int32(localGame.Tableau[i].HiddenCount)
		cardsLen := len(localGame.Tableau[i].Cards)
		pbGame.Stack[i].Cards = make([]*pb.Card, cardsLen)
		for j := 0; j < cardsLen; j++ {
			localCard := localGame.Tableau[i].Cards[j]
			pbGame.Stack[i].Cards[j] =
				&pb.Card{
					Suit: int32(localCard.Suit),
					Rank: int32(localCard.Rank),
				}
		}
	}

	s.active[pbGame.Id] = localGame
	return &pbGame, nil
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
