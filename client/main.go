package main

import (
	"context"
	"fmt"
	"log"
	"time"

	service "grpc-project/client/client_grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close() // Close connection when done

	// Create the Greeter client for SayHello (hello.proto)
	greeterClient := service.NewGreeterClient(conn)

	// Call SayHello function
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	helloReq := &service.HelloRequest{Name: "Shofwan Shiddiq"}
	helloRes, err := greeterClient.SayHello(ctx, helloReq)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	// Print result for HelloRequest
	fmt.Println("Greeting from SayHello:", helloRes.GetMessage())

	// Create the UserService client for GetUser (user.proto)
	userServiceClient := service.NewUserServiceClient(conn)

	// Call GetUser function
	userReq := &service.UserRequest{UserID: 1}
	userRes, err := userServiceClient.GetUser(ctx, userReq)
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	// Print result for UserRequest
	fmt.Printf("User Info: ID=%d, Name=%s, Email=%s\n", userRes.GetUserID(), userRes.GetName(), userRes.GetEmail())
}
