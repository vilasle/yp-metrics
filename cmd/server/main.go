package main

import (
	"fmt"

	"github.com/vilasle/yp-metrics/internal/repository/memory"
	"github.com/vilasle/yp-metrics/internal/service"
	"github.com/vilasle/yp-metrics/internal/transport/rest"
)

// http://localhost:8080/update/<metrisType>/<metricName>/<metricValue>
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("application is in a panic", err)
		}
	}()

	gaugeStorage := memory.NewMetricGaugeMemoryRepository()
	counterStorage := memory.NewMetricCounterMemoryRepository()

	svc := service.NewStorageService(gaugeStorage, counterStorage)

	server := rest.NewHttpServer(":8080")

	server.Register("/update/", rest.UpdateHandler(svc))

	if err := server.Start(); err != nil {
		fmt.Printf("server starting got error, %v", err)
	}

}
