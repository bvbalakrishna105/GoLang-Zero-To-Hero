package main

import (
	"context"
	"log"
	"net"

	pb "github.com/bvbalakrishna105/GoLang-Zero-To-Hero/venky-grpc-go/chat" // Import generated code
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer // Embed the generated struct
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello, " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Println("Server started at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
