package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure" 
	pb "pinger/gencode"                       
)



var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to") // Command-line flag to specify the server address.
)

func main() {
	flag.Parse() 

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close() 

	c := pb.NewPingServiceClient(conn)

	
	pingCtx, cancelPing := context.WithTimeout(context.Background(), time.Second) // Create a context with a timeout for ping.
	defer cancelPing()                                                             // Cancel the ping context when main function exits.

	pingResponse, err := c.HealthCheck(pingCtx, &pb.PingRequest{})
	if err != nil {
		log.Fatalf("could not ping server: %v", err)
	}
	log.Printf("Server response to ping: %s", pingResponse.GetMessage()) // Log the ping response message.
}
