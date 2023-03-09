package server

import (
	"net"

	"github.com/dstopka/notebook-app/backend/common/genproto/users"
	"google.golang.org/grpc"
)

func RunGRPCServer(addr string, srv users.UsersServiceServer) error {
	server := grpc.NewServer()
	users.RegisterUsersServiceServer(server, srv)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return server.Serve(listener)
}