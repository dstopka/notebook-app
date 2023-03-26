package infrastructure

import (
	"context"

	"github.com/dstopka/notebook-app/backend/common/genproto/notebooks"
	"github.com/dstopka/notebook-app/backend/notebooks/internal/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ notebooks.NotebooksServiceServer = (*GrpcServer)(nil)

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

// CreateNotebook implements NotebooksServiceServer's CreateNotebook method.
func (g *GrpcServer) CreateNotebook(ctx context.Context, req *notebooks.CreateNotebookRequest) (*notebooks.Notebook, error) {
	return nil, status.Errorf(codes.Unimplemented, "CreateNotebook not implemented")
}

// ListNotebooks implements NotebooksServiceServer's ListNotebooks method.
func (g *GrpcServer) ListNotebooks(ctx context.Context, req *notebooks.ListNotebooksRequest) (*notebooks.Notebooks, error) {
	return nil, status.Errorf(codes.Unimplemented, "ListNotebooks not implemented")
}

// GetNotebook implements NotebooksServiceServer's GetNotebook method.
func (g *GrpcServer) GetNotebook(ctx context.Context, req *notebooks.GetNotebookRequest) (*notebooks.Notebook, error) {
	return nil, status.Errorf(codes.Unimplemented, "GetNotebook not implemented")
}

// UpdateNotebook implements NotebooksServiceServer's UpdateNotebook method.
func (g *GrpcServer) UpdateNotebook(ctx context.Context, req *notebooks.UpdateNotebookRequest) (*notebooks.Notebook, error) {
	return nil, status.Errorf(codes.Unimplemented, "UpdateNotebook not implemented")
}

// DeleteNotebook implements NotebooksServiceServer's DeleteNotebook method.
func (g *GrpcServer) DeleteNotebook(ctx context.Context, req *notebooks.DeleteNotebookRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "DeleteNotebook not implemented")
}