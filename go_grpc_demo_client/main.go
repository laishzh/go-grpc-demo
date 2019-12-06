package main

import (
	"context"
	"fmt"
	pb "github.com/laishzh/go-grpc-demo/pb"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

var (
	address = "localhost:50051"
)

func init() {
	if len(os.Args) > 1 {
		address = os.Args[1]
	}
}

func main() {
	fmt.Println("Client Start!!")
	fmt.Printf("Connect to: %s\n", address)

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
