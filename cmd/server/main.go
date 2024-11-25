package main

import (
	"fmt"

	"github.com/vilasle/yp-metrics/internal/repository/memory"
	svc "github.com/vilasle/yp-metrics/internal/service/server"
	rest "github.com/vilasle/yp-metrics/internal/transport/rest/server"
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

	svc := svc.NewStorageService(gaugeStorage, counterStorage)

	server := rest.NewHTTPServer(":8080")

	server.Register("/update/", rest.UpdateHandler(svc))

	if err := server.Start(); err != nil {
		fmt.Printf("server starting got error, %v", err)
	}

}
