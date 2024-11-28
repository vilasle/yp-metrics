package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"math/rand"

	"github.com/vilasle/yp-metrics/internal/service/agent/collector"
	"github.com/vilasle/yp-metrics/internal/service/agent/sender/rest"
)

func main() {

	endpoint := flag.String("a", "localhost:8080", "endpoint to send metrics")
	reportSec := flag.Int("r", 10, "timeout(sec) for sending report to server")
	pollSec := flag.Int("p", 2, "timeout(sec) for polling metrics")

	flag.Parse()

	c := collector.NewRuntimeCollector()

	metrics := defaultGaugeMetrics()
	err := c.RegisterMetric(metrics...)

	if err != nil {
		fmt.Printf("can to register metric by reason %v\n", err)
		os.Exit(1)
	}

	pollInterval := time.Second * time.Duration(*pollSec)
	reportInterval := time.Second * time.Duration(*reportSec)

	c.RegisterEvent(func(c *collector.RuntimeCollector) {
		counter := c.GetCounterValue("PullCount")
		counter.Increment()

		c.SetCounterValue(counter)
	})

	c.RegisterEvent(func(c *collector.RuntimeCollector) {
		gauge := c.GetGaugeValue("RandomValue")
		gauge.SetValue(rand.Float64())

		c.SetGaugeValue(gauge)

	})

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	pollTicker := time.NewTicker(pollInterval)
	reportTicker := time.NewTicker(reportInterval)

	updateAddress := fmt.Sprintf("http://%s/update/", *endpoint)

	fmt.Printf("sending metrics to %s\n", updateAddress)
	fmt.Printf("polling metrics every %d sec\n", *pollSec)
	fmt.Printf("sending report every %d sec\n", *reportSec)

	fmt.Println("press ctrl+c to exit")

	sender, err := rest.NewHTTPSender(updateAddress)
	if err != nil {
		fmt.Printf("can not create sender by reason %v", err)
		os.Exit(2)
	}

	agent := NewCollectorAgent(c, sender)

	for run := true; run; {
		select {
		case <-pollTicker.C:
			agent.Collect()
		case <-reportTicker.C:
			agent.Report()
		case <-sigint:
			pollTicker.Stop()
			reportTicker.Stop()
			run = false
		}
	}
}

func defaultGaugeMetrics() []string {
	return []string{
		"Alloc",
		"BuckHashSys",
		"Frees",
		"GCCPUFraction",
		"GCSys",
		"HeapAlloc",
		"HeapIdle",
		"HeapInuse",
		"HeapObjects",
		"HeapReleased",
		"HeapSys",
		"LastGC",
		"Lookups",
		"MCacheInuse",
		"MCacheSys",
		"MSpanInuse",
		"MSpanSys",
		"Mallocs",
		"NextGC",
		"NumForcedGC",
		"NumGC",
		"OtherSys",
		"PauseTotalNs",
		"StackInuse",
		"StackSys",
		"Sys",
		"TotalAlloc",
	}
}
