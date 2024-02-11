package main

import (
	"context"
	"log"

	pb "your_module_path/your_service" // Import generated code

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewYourServiceClient(conn)

	// Call the service method
	response, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %v", err)
	}
	log.Printf("Response from server: %s", response.Message)
}
