package server

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func RunGRPCServer(port string, registerServer func(server *grpc.Server)) {
	if port == "" {
		port = "8080"
	}

	flag.Parse()

	addrs := fmt.Sprintf(":%s", port)
	RunGRPCServerInAddrs(addrs,
		registerServer)
}

func RunGRPCServerInAddrs(addrs string, registerServer func(server *grpc.Server)) {
	lis, err := net.Listen("tcp", addrs)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	registerServer(grpcServer)
	grpcServer.Serve(lis)

}
