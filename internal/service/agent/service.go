package agent

import "github.com/vilasle/yp-metrics/internal/metric"

type Collector interface {
	Collect()
	AllMetrics() []metric.Metric
}

type Sender interface {
	Send(metric.Metric) (error)
}