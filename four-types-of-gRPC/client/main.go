package main

import (
	"log"

	pb "github.com/anik-adnan/grpc-demo/protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	connection, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}
	defer connection.Close()

	client := pb.NewGreeterServiceClient(connection)

	names := &pb.Namelist{
		Names: []string{"Anik", "Adnan", "Shamim", "Gazi noor"},
	}
	// callSayHello(client)
	// callSayHelloServiceStream(client, names)
	// callSayHelloClientStream(client, names)
	callSayHelloBidirectionalStreaming(client, names)

}
