package infrastructure

import (
	"context"

	"github.com/dstopka/notebook-app/backend/common/genproto/users"
	"github.com/dstopka/notebook-app/backend/users/internal/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GrpcServer is the implementation of UserServiceServer.
type GrpcServer struct {
	app *app.Application
	users.UnimplementedUsersServiceServer
}

// NewGrrpcServer returns a new GrpcServer.
func NewGrpcServer(app *app.Application) *GrpcServer {
	if app == nil {
		panic("app is nil")
	}

	return &GrpcServer{
		app: app,
	}
}

// GetUser implements UserServiceServer's GetUser method.
func (g *GrpcServer) GetUser(ctx context.Context, req *users.GetUserRequest) (*users.GetUserResponse, error) {
	user, err := g.app.HandleGetUser(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &users.GetUserResponse{
		DisplayName: user.Name,
		AvatarUrl:   user.AvatarURL,
		Role:        user.Role,
	}

	return resp, nil
}

// SetAvatar implements UserServiceServer's SetAvatar method.
func (g *GrpcServer) SetAvatar(ctx context.Context, req *users.SetAvatarRequest) (*emptypb.Empty, error) {
	if err := g.app.HandleSetAvatar(ctx, req.UserId, req.AvatarUrl); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

// SetLastIP implements UserServiceServer's SetLastIP method.
func (g *GrpcServer) SetLastIP(ctx context.Context, req *users.SetLastIPRequest) (*emptypb.Empty, error) {
	if err := g.app.HandleSetLastIP(ctx, req.UserId, req.Ip); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
