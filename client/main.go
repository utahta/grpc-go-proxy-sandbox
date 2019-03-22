package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/utahta/grpc-go-proxy-example/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	address     = "localhost:50052"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Outgoing: abe hiroshi")
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("abe", "hiroshi"))

	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
