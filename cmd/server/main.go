package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

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

	stop := make(chan os.Signal, 1)
	defer close(stop)

	signal.Notify(stop, os.Interrupt)

	go func() {
		if err := server.Start(); err != nil {
			fmt.Printf("server starting got error, %v", err)
		}
		stop <- os.Interrupt
	}()

	<-stop

	if !server.IsRunning() {
		os.Exit(0)
	}

	stopErr := make(chan error)
	defer close(stopErr)

	tickForce := time.NewTicker(time.Second * 5)
	tickKill := time.NewTicker(time.Second * 10)

	go func() { stopErr <- server.Stop() }()

	for {
		select {
		case err := <-stopErr:
			if err != nil {
				fmt.Println("server stopped with error", err)
			}
			server.ForceStop()
		case <-tickForce.C:
			go server.ForceStop()
		case <-tickKill.C:
			fmt.Println("server did not stop during expected time")
			os.Exit(1)
		}
	}

}
