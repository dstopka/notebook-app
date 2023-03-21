package server

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

// GRPCServerConfig defines all methods needed to configure a grpc service server.
type GRPCServerConfig interface {
	Addr() string
}

// RegisterServerFn defines function used to register grpc server.
type RegisterServerFn func(*grpc.Server)

func RunGRPCServer(c GRPCServerConfig, registerServer RegisterServerFn) error {
	server := grpc.NewServer()
	registerServer(server)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func() {
		<-ctx.Done()
		server.GracefulStop()
	}()

	listener, err := net.Listen("tcp", c.Addr())
	if err != nil {
		return err
	}
	return server.Serve(listener)
}
