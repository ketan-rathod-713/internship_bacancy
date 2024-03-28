package main

import (
	"context"
	pb "grpcproject/proto"
)

// unary serveer function
func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
