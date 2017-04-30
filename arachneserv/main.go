package main

import (
	//	"context"
	"fmt"
	"net"

	context "golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/dougfort/arachne/arachne"
)

const port = 10000

type arachneServer struct {
}

// StartGame starts a new game
func (s *arachneServer) StartGame(
	ctx context.Context,
	request *pb.GameRequest,
) (*pb.Game, error) {
	return &pb.Game{Id: 666, Seed: 42}, nil
}

func newServer() *arachneServer {
	s := new(arachneServer)
	return s
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
