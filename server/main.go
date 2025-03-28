package main

import (
	"context"
	"fmt"
	service "grpc-project/server/server_grpc"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Inititate GRPC Server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Initiate
	grpcServer := grpc.NewServer()

	// Service Registry
	service.RegisterGreeterServer(grpcServer, &server{})
	service.RegisterUserServiceServer(grpcServer, &server{})

	// accept gRPC server reflection
	reflection.Register(grpcServer)

	fmt.Println("GRPC Server started on port 50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

type server struct {
	service.UnimplementedGreeterServer
	service.UnimplementedUserServiceServer
}

func (s *server) SayHello(
	ctx context.Context, //
	req *service.HelloRequest,
) (*service.HelloResponse, error) {
	return &service.HelloResponse{
		Message: "Hello " + req.Name,
	}, nil
}

func (s *server) GetUser(
	ctx context.Context,
	req *service.UserRequest,
) (*service.UserResponse, error) {
	userData := map[int32]*service.UserResponse{
		1:  {UserID: 1, Name: "Erling Haaland", Email: "haaland@manchestercity.com"},
		2:  {UserID: 2, Name: "Kevin De Bruyne", Email: "debruyne@manchestercity.com"},
		3:  {UserID: 3, Name: "Riyad Mahrez", Email: "mahrez@manchestercity.com"},
		4:  {UserID: 4, Name: "Jack Grealish", Email: "grealish@manchestercity.com"},
		5:  {UserID: 5, Name: "Bernardo Silva", Email: "silva@manchestercity.com"},
		6:  {UserID: 6, Name: "Phil Foden", Email: "foden@manchestercity.com"},
		7:  {UserID: 7, Name: "Rodri", Email: "rodri@manchestercity.com"},
		8:  {UserID: 8, Name: "Ilkay Gundogan", Email: "gundogan@manchestercity.com"},
		9:  {UserID: 9, Name: "Joao Cancelo", Email: "cancelo@manchestercity.com"},
		10: {UserID: 10, Name: "Kyle Walker", Email: "walker@manchestercity.com"},
	}

	if user, ok := userData[req.UserID]; ok {
		return user, nil
	}

	return &service.UserResponse{}, nil
}
