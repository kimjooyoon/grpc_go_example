package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	hello "grpc_go_example/example/helloworld"
	"log"
	"net"
)

const (
	defaultName = "world"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	hello.UnimplementedGreeterServer
}

func (s *server) SayHello(_ context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &hello.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err1 := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err1 != nil {
		log.Fatalf("failed to listen: %v", err1)
	}
	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err2 := s.Serve(lis); err2 != nil {
		log.Fatalf("failed to serve: %v", err2)
	}

}
