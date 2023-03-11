package infrastructure

import (
	"context"

	"github.com/dstopka/notebook-app/backend/common/genproto/users"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	users.UnimplementedUsersServiceServer
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

func (g *GrpcServer) GetUser(_ context.Context, _ *users.GetUserRequest) (*users.GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "GetUser not implemented")
}

func (g *GrpcServer) SetAvatar(_ context.Context, _ *users.SetAvatarRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "SetAvatar not implemented")
}

func (g *GrpcServer) SetLastIP(_ context.Context, _ *users.SetLastIPRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "SetLastIP not implemented")
}