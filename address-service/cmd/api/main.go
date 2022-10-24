package main

import (
	"address-service/cmd/initializers"
	"address-service/cmd/mapper"
	"address-service/cmd/proto"
	"address-service/cmd/pubsub"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type AddressServer struct {
	proto.UnimplementedAddressRegistrationServer
}

func (u *AddressServer) RegisterAddress(ctx context.Context, req *proto.AddressRequest) (*proto.AddressResponse, error) {

	model := mapper.ProtoToModel(req)
	result := initializers.DB.Create(&model)
	if result.Error != nil {
		log.Fatalf("address could not be registered: %v", result.Error)
		return nil, result.Error
	}
	response := mapper.ModelToProto(&model)
	log.Println("Address registered")

	return &response, nil
}

func init() {
	initializers.Migrate()
}

func main() {
	listener, err := net.Listen("tcp", ":82")
	if err != nil {
		log.Fatalf("could not start listening: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterAddressRegistrationServer(s, &AddressServer{})

	reflection.Register(s)

	log.Println("Server started running on port 82")

	log.Println("Started listening for the incoming pubsub messages")

	go pubsub.Connect()

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("server could not start: %v", err)
	}

}
