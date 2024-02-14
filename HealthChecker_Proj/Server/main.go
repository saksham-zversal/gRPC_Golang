package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	pb "pinger/gencode"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedPingServiceServer
}

func (s *server) HealthCheck(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	log.Println("Received ping from client")
	return &pb.PingReply{Message: "Health Check Successfull"}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	s := grpc.NewServer()

	pb.RegisterPingServiceServer(s, &server{})

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
