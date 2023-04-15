package server

import (
	"flag"
	"fmt"
	"log"
	"net"

	"golang.org/x/exp/slog"
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

func RunGRPCServerInAddrs(
	addrs string,
	registerServer func(server *grpc.Server),
) {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	registerServer(grpcServer)

	lis, err := net.Listen("tcp", addrs)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	slog.Info("Starting gRPC server on port", slog.String("port", addrs))

	log.Fatal(grpcServer.Serve(lis))
}
