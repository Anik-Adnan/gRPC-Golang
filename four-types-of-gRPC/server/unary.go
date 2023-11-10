package main

import (
	"context"

	pb "github.com/anik-adnan/grpc-demo/protoc"
)

func (s *helloServer) Sayhello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello from unary response",
	}, nil
}
