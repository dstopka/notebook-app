package infrastructure

import (
	"github.com/dstopka/notebook-app/backend/common/genproto/users"
)

type GrpcServer struct {
	users.UnimplementedUsersServiceServer
}