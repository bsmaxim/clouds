package main

import (
	"context"
	pb "hello_grpc/helloworld"
	"log"

	"hello_grpc/helloworld"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s server) SayHello(ctx context.Context, in *helloworld.HelloRequest, opts ...grpc.CallOption) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
