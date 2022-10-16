package main

import (
	"fmt"

	"github.com/blockami/services/information/handler"
	pb "github.com/blockami/services/information/proto"

	"github.com/micro/micro/plugin/prometheus/v3"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/metrics"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("information"),
	)

	reporter, err := prometheus.New(metrics.Address(":9002"))
	if err != nil {
		fmt.Printf("Error setting reporter: %v", err)
	}
	metrics.SetDefaultMetricsReporter(reporter)

	// Register handler
	pb.RegisterInformationHandler(srv.Server(), handler.New(srv))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
