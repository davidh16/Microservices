package main

import (
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
	"user-registration/cmd/address-service"
	"user-registration/cmd/initializers"
	"user-registration/cmd/mapper"
	"user-registration/cmd/proto"
	"user-registration/cmd/pubsub"
	"user-registration/cmd/user-service"
	"user-registration/cmd/utils"
)

type UserRegistrationServer struct {
	proto.UnimplementedUserRegistrationServer
}

func (u *UserRegistrationServer) RegisterUser(ctx context.Context, req *proto.RegistrationRequest) (*proto.RegistrationResponse, error) {
	userRequest := mapper.RegReqToUsrReq(req)
	userResponse, err := user_service.CreateUser(userRequest)
	if err != nil {
		log.Fatalf("could not create user: %v", err)
		return nil, err
	}

	addressRequest := mapper.RegReqToAdrReq(req)
	addressResponse, err := address_service.RegisterAddress(&addressRequest)
	if err != nil {
		log.Fatalf("could not register address: %v", err)
	}

	if userResponse != nil && addressResponse != nil {
		userToRegister := mapper.ProtoToModel(req)
		err2, registeredUser := utils.SaveUser(userToRegister)
		if err2 != nil {
			log.Fatalf("user could not be saved to database: %v", err)
			return nil, err2
		}

		id := strconv.FormatInt(userResponse.Id, 10)

		pubsub.PublishCommit(id)

		message, _ := json.Marshal(registeredUser)

		pubsub.PublishUser(message)

		response := proto.RegistrationResponse{
			Id:        userResponse.Id,
			Name:      userResponse.Name,
			Surname:   userResponse.Surname,
			AddressId: addressResponse.Id,
			Address:   addressResponse.Address,
		}
		return &response, nil
	}

	err = errors.New("unexpected error")
	return nil, err

}

func init() {
	initializers.Migrate()
}

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("could not start listening: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterUserRegistrationServer(s, &UserRegistrationServer{})

	reflection.Register(s)

	log.Println("Server started running on port 80")

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("server could not start: %v", err)
	}

}
