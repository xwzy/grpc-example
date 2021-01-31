package main

import (
	"context"

	"github.com/xwzy/grpc-example/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":50051"

type server struct {
	hello.UnimplementedHelloServer
}

func (s *server) GetGreeting(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &hello.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	hello.RegisterHelloServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
