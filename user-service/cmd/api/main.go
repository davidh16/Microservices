package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"user-service/cmd/initializers"
	"user-service/cmd/mapper"
	"user-service/cmd/proto"
	"user-service/cmd/pubsub"
)

type UserServer struct {
	proto.UnimplementedUserServer
}

func (u *UserServer) RegisterUser(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {

	model := mapper.ProtoToModel(req)
	result := initializers.DB.Create(&model)
	if result.Error != nil {
		log.Fatalf("user could not be created: %v", result.Error)
		return nil, result.Error
	}
	response := mapper.ModelToProto(&model)
	log.Println("User created")

	return &response, nil
}

func init() {
	initializers.Migrate()
}

func main() {
	listener, err := net.Listen("tcp", ":81")
	if err != nil {
		log.Fatalf("could not start listening: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterUserServer(s, &UserServer{})

	reflection.Register(s)

	log.Println("Server started running on port 81")

	log.Println("Started listening for the incoming pubsub messages")

	go pubsub.Connect()

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("server could not start: %v", err)
	}
}
