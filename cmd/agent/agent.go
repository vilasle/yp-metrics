package main

import (
	"sync"

	"github.com/vilasle/yp-metrics/internal/service/agent"
)

type collectorAgent struct {
	agent.Collector
	agent.Sender
	mx *sync.Mutex
}

func NewCollectorAgent(collector agent.Collector, sender agent.Sender) collectorAgent {
	return collectorAgent{
		Collector: collector,
		Sender:    sender,
		mx:        &sync.Mutex{},
	}
}

func (a collectorAgent) Collect() {
	a.mx.Lock()
	defer a.mx.Unlock()

	a.Collector.Collect()
}

func (a collectorAgent) Report() {
	a.mx.Lock()
	defer a.mx.Unlock()

	for _, metric := range a.Collector.AllMetrics() {
		a.Sender.Send(metric)
	}
}
