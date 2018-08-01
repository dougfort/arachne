package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/ardanlabs/kit/cfg"

	pb "github.com/dougfort/arachne/arachne"
)

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

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			grpclog.Errorf("grpcServer.Serve failed: %s", err)
		}
	}()

	s := <-sigChan
	grpclog.Infof("server received %v: terminating", s)
	grpcServer.Stop()
}
