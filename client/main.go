package main

import (
	"context"
	"fmt"
	pb "github.com/laishzh/go-grpc-demo/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	fmt.Println("Client Start!!")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	r, err := c.SayHello(ctx, &pb.HelloRequest{
		Name: "Name",
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())
}
