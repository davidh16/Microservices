package address_service

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
	"user-registration/cmd/proto/address-service"
)

func RegisterAddress(req *proto.AddressRequest) (*proto.AddressResponse, error) {
	conn, err := grpc.Dial("address-service:82", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Printf("could not connect to register service: %v", err)
		return nil, err
	}
	//defer conn.Close()
	c := proto.NewAddressRegistrationClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	addressRegistrationResponse, err := c.RegisterAddress(ctx, req)
	if err != nil {
		log.Printf("could not register user: %v", err)
		return nil, err
	}

	return addressRegistrationResponse, nil
}
