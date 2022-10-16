package main

import (
	"github.com/blockami/servicestoken/handler"
	pb "github.com/blockami/servicestoken/proto"

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
