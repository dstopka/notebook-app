package main

import (
	"fmt"
	"log"

	"github.com/dstopka/notebook-app/backend/users/internal/infrastructure"
	"github.com/dstopka/notebook-app/backend/users/pkg/server"
)

func main() {
	port := "8080"
	addr := fmt.Sprintf(":%s", port)
	
	srv := infrastructure.GrpcServer{}

	if err := server.RunGRPCServer(addr, srv); err != nil {
		log.Fatalln(err)
	}
}