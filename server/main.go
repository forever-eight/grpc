package main

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/forever-eight/grpc.git/proto"
)

func main() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterReverseServer(grpcServer, &server{})

	err = grpcServer.Serve(listener)
	if err != nil {
		grpclog.Fatalf("failed to serve: %v", err)
	}

}

type server struct {
	pb.UnimplementedReverseServer
}

func (s *server) Do(c context.Context, request *pb.ReverseRequest) (response *pb.ReverseResponse, err error) {
	n := 0
	// Create an array of runes to safely reverse a string.
	rune := make([]rune, len(request.Message))

	for _, r := range request.Message {
		rune[n] = r
		n++
	}

	// Reverse using runes.
	rune = rune[0:n]

	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}

	output := string(rune)
	response = &pb.ReverseResponse{
		Message: output,
	}

	return response, nil
}
