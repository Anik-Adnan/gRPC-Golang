package main

import (
	"log"
	"net"

	pb "github.com/anik-adnan/grpc-demo/protoc"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreeterServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServiceServer(grpcServer, &helloServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("could not greet: %v", err)
	}

}
