package main

import (
	"github.com/blockami/pricing/handler"
	pb "github.com/blockami/pricing/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("pricing"),
	)

	// Register handler
	pb.RegisterPricingHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
