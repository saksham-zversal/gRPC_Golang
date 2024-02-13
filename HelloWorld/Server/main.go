package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "helloworld/gencode"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server implements the Greeter service.
type server struct {
	pb.UnimplementedGreeterServer // Embed the generated UnimplementedGreeterServer to implement the GreeterServer interface.
}

// SayHello implements the SayHello RPC method of the Greeter service.
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil // Respond with a greeting message.
}

func main() {
	flag.Parse()

	// Create a listener on the specified port.
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close() // Close the listener when main function exits.

	// Create a new gRPC server.
	s := grpc.NewServer()

	// Register the Greeter service with the gRPC server.
	pb.RegisterGreeterServer(s, &server{})

	// Start serving on the listener.
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
