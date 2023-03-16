package server

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/dstopka/notebook-app/backend/common/genproto/users"
	"google.golang.org/grpc"
)

func RunGRPCServer(addr string, srv users.UsersServiceServer) error {
	server := grpc.NewServer()
	users.RegisterUsersServiceServer(server, srv)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func() {
		<-ctx.Done()
		server.GracefulStop()
	}()

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return server.Serve(listener)
}
