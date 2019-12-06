//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a go_grpc_demo_server for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/laishzh/go-grpc-demo/pb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// go_grpc_demo_server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

var name = "NoOne"

func init() {
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v\n", in.GetName())
	msg := fmt.Sprintf("Message from %s", name)
	return &pb.HelloReply{Message: msg}, nil
}

func main() {

	fmt.Printf("Server is starting, go_grpc_demo_server name: %s\n", name)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
