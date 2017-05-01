package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/dougfort/arachne/arachne"
)

const serverAddr = "127.0.0.1:10000"

func newGame() (*pb.Game, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewArachneClient(conn)

	gm, err := client.StartGame(context.Background(), &pb.GameRequest{})
	if err != nil {
		grpclog.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}

	return gm, nil
}
