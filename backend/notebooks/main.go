package main

import (
	"fmt"
	"log"

	"github.com/dstopka/notebook-app/backend/common/server"
	"github.com/dstopka/notebook-app/backend/notebooks/internal/config"
	"github.com/dstopka/notebook-app/backend/notebooks/internal/infrastructure"
	"github.com/dstopka/notebook-app/backend/notebooks/internal/service"
	"github.com/dstopka/notebook-app/backend/notebooks/pkg/notebooks"
	"google.golang.org/grpc"
)

func main() {
	port := "8080"
	addr := fmt.Sprintf(":%s", port)
	config := config.Config{Address: addr}

	app := service.NewApplication()

	err := server.RunGRPCServer(config, func(s *grpc.Server) {
		srv := infrastructure.NewGrpcServer(app)
		notebooks.RegisterNotebooksServiceServer(s, srv)
	})
	if err != nil {
		log.Fatalln(err)
	}
}
