package main

import (
	"fmt"

	"github.com/blockami/services/transactions/handler"
	pb "github.com/blockami/services/transactions/proto"

	"github.com/micro/micro/plugin/prometheus/v3"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/metrics"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("transactions"),
	)
	srv.Init()

	reporter, err := prometheus.New(metrics.Address(":9000"))
	if err != nil {
		fmt.Printf("Error setting reporter: %v", err)
	}
	metrics.SetDefaultMetricsReporter(reporter)

	// Register handlerw
	pb.RegisterTransactionsHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
