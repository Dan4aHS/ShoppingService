package grpcapp

import (
	"ShoppingExpensesService/internal/config"
	"ShoppingExpensesService/internal/service"
	"ShoppingExpensesService/internal/transport/grpchandlers"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func NewApp(ps service.IProductService, cfg *config.Config) *App {
	grpcServer := grpc.NewServer()
	grpchandlers.RegisterGRPCServer(grpcServer, ps, cfg)
	return &App{
		gRPCServer: grpcServer,
		port:       cfg.App.GRPC.Port,
	}
}

func (a *App) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("grpc server listening on port %d", a.port)
	err = a.gRPCServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
