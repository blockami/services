package main

import (
	"github.com/blockami/services/token/handler"
	pb "github.com/blockami/services/token/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("token"),
	)

	// Register handler
	pb.RegisterTokenHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
