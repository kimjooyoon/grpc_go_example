package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	hello "grpc_go_example/example/helloworld"
	"log"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	conn, err1 := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err1 != nil {
		log.Fatalf("did not connect: %v", err1)
	}
	defer conn.Close()
	c := hello.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err2 := c.SayHello(ctx, &hello.HelloRequest{Name: *name})
	if err2 != nil {
		log.Fatalf("could not greet: %v", err2)
	}
	log.Printf("Greeting: %s", r.GetMessage())

}
