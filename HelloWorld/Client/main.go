package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure" // Importing insecure credentials for development purposes.
	pb "helloworld/gencode"                       // Importing the generated gRPC code package.
)

const (
	defaultName = "world" // Default name to greet.
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to") // Command-line flag to specify the server address.
	name = flag.String("name", defaultName, "Name to greet")                   // Command-line flag to specify the name to greet.
)

func main() {
	flag.Parse() // Parsing command-line flags.
	// Set up a connection to the server with insecure credentials.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close() // Close the connection when main function exits.

	// Create a new gRPC client.
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // Create a context with a timeout.
	defer cancel()                                                        // Cancel the context when main function exits.
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})              // Invoke the SayHello RPC with the provided name.
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage()) // Log the response message.
}
