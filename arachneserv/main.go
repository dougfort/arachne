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
	var gm pb.Game

	gm.Id = s.nextID
	s.nextID++

	switch request.Gametype {
	case pb.GameRequest_RANDOM:
		s.active[gm.Id] = game.New()
	case pb.GameRequest_REPLAY:
		s.active[gm.Id] = game.Replay(request.Seed)
	default:
		return nil, fmt.Errorf("invalid GameType %d", request.Gametype)
	}

	gm.Seed = s.active[gm.Id].Deck.Seed()

	return &gm, nil
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
