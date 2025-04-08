package grpccon

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func NewGrpcServer(host string) (*grpc.Server, net.Listener) {
	opts := make([]grpc.ServerOption, 0)
	grpcServer := grpc.NewServer(opts...)
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Error: Failed to listen: %v", err)
	}
	return grpcServer, lis
}
