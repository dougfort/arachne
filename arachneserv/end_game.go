package main

import (
	//	"context"
	context "golang.org/x/net/context"

	pb "github.com/dougfort/arachne/arachne"
)

// EndGame requests an end to the game
func (s *arachneServer) EndGame(
	ctx context.Context,
	request *pb.EndGameRequest,
) (*pb.EndGameResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.active, request.Id)

	return &pb.EndGameResponse{}, nil
}
