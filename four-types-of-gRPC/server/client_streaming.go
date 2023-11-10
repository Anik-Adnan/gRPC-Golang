package main

import (
	"io"
	"log"

	pb "github.com/anik-adnan/grpc-demo/protoc"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreeterService_SayHelloClientStreamingServer) error {
	var meassages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: meassages})
		}
		if err != nil {
			return err
		}
		log.Printf("got request with name: %v", req.Name)
		meassages = append(meassages, " hello ", req.Name)
	}
}
