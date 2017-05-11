package main

import (
	//	"context"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/dougfort/arachne/game"

	pb "github.com/dougfort/arachne/arachne"
)

type arachneServer struct {
	mutex  sync.Mutex
	nextID int64
	active map[int64]*game.Game
}

func newServer() *arachneServer {
	var s arachneServer
	s.nextID = 1
	s.active = make(map[int64]*game.Game)
	return &s
}

func main() {
	address := getAddressFromEnv()

	grpclog.Printf("listening to: %s", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterArachneServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
