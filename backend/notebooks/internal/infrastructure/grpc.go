package infrastructure

import (
	"github.com/dstopka/notebook-app/backend/common/genproto/notebooks"
	"github.com/dstopka/notebook-app/backend/notebooks/internal/app"
)

var _ notebooks.UnsafeNotebooksServiceServer = (*GrpcServer)(nil)

// GrpcServer defines the implementation of NotebooksServiceServer.
type GrpcServer struct {
	app *app.Application
	notebooks.UnsafeNotebooksServiceServer
}

// NewGrpcServer creates a new GrpcServer using provided app.
func NewGrpcServer(app *app.Application) *GrpcServer {
	if app == nil {
		panic("app is nil")
	}

	return &GrpcServer{
		app: app,
	}
}