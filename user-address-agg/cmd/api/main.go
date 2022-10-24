package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"user-address-agg/cmd/initializers"
	"user-address-agg/cmd/proto"
	"user-address-agg/cmd/pubsub"
)

type UserAddressAggServer struct {
	proto.UnimplementedAggregatorServer
}

func init() {
	initializers.ConnectToMongo()
}

func main() {
	listener, err := net.Listen("tcp", ":83")
	if err != nil {
		log.Fatalf("could not start listening: %v", err)
	}

	s := grpc.NewServer()

	proto.RegisterAggregatorServer(s, &UserAddressAggServer{})

	reflection.Register(s)

	log.Println("Server started on port 83")

	log.Println("Started listening for the incoming pubsub messages")

	go pubsub.Connect()

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("server could not start: %v", err)
	}
}
