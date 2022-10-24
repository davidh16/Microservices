package user_service

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
	"user-registration/cmd/proto/user-service"
)

func CreateUser(req *user_service.UserRequest) (*user_service.UserResponse, error) {
	conn, err := grpc.Dial("user-service:81", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Printf("could not connect to register service: %v", err)
		return nil, err
	}
	defer conn.Close()
	c := user_service.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	userRegistrationResponse, err := c.RegisterUser(ctx, req)
	if err != nil {
		log.Printf("could not register user: %v", err)
		return nil, err
	}

	return userRegistrationResponse, nil
}
