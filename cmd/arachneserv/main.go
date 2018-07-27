package main

import (
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/ardanlabs/kit/cfg"

	"github.com/dougfort/arachne/internal/game"

	pb "github.com/dougfort/arachne/arachne"
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

func main() {
	const cfgNamespace = "arachne"
	var address string

	if err := cfg.Init(cfg.EnvProvider{Namespace: cfgNamespace}); err != nil {
		panic(err)
	}

	address = cfg.MustString("ADDRESS")

	grpclog.Infof("listening to: %s", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterArachneServer(grpcServer, newServer())
	if err := grpcServer.Serve(lis); err != nil {
		grpclog.Errorf("grpcServer.Serve failed: %s", err)
	}
}
