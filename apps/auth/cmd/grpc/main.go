package main

import (
	"os"

	"github.com/eduaravila/momo/apps/auth/internal/ports"
	"github.com/eduaravila/momo/apps/auth/internal/service"
	"github.com/eduaravila/momo/packages/genproto/auth"
	"github.com/eduaravila/momo/packages/server"
	"google.golang.org/grpc"
)

func main() {
	grcpPort := os.Getenv("AUTH_GRPC_PORT")
	if grcpPort == "" {
		grcpPort = "3002"
	}

	app := service.NewApplication()

	go server.RunGRPCServer(":"+grcpPort, func(server *grpc.Server) {
		auth.RegisterSessionServiceServer(server, ports.NewGRPCServer(app))
	})
}
