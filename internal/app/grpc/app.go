package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	userGrpcSrv "user-service/internal/grpc/user"
	userService "user-service/internal/services/user"
)

type App struct {
	grpcServer *grpc.Server
	port       int
}

func New(userService *userService.UserService, port int) *App {
	grpcServer := grpc.NewServer()
	userGrpcSrv.Register(grpcServer, userService)

	return &App{
		grpcServer: grpcServer,
		port:       port,
	}
}

func (a *App) Serve() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return err
	}

	log.Println("gRPC server running on", l.Addr())

	if err := a.grpcServer.Serve(l); err != nil {
		return err
	}

	return nil
}

func (a *App) Stop() {
	log.Println("gRPC server stopping")
	a.grpcServer.GracefulStop()
}
