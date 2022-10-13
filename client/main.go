package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const defaultName = "world"

var (
	connectionAddress = flag.String("addr", "localhost:50051", "Connection Address")
	greetingName      = flag.String("name", defaultName, "Greeting Name")
)

func main() {
	flag.Parse()

	connection, err := grpc.Dial(
		*connectionAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("Connection Error %v", err)
	}

	defer connection.Close()

	greetingClient := pb.NewGreeterClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	response, err := greetingClient.SayHello(ctx, &pb.HelloRequest{Name: *greetingName})

	if err != nil {
		log.Fatalf("Error Response %v", err)
	}

	log.Printf("Greeting: %s", response.GetMessage())
}
