package service

import "github.com/vilasle/yp-metrics/internal/metric"

//server interfaces

type StorageService interface {
	Save(metric.RawMetric) error
	Get(string) (metric.Metric, error)
	AllMetrics() []metric.Metric
}

//agent interfaces

type Collector interface {
	Collect()
	AllMetrics() []metric.Metric
}

type Sender interface {
	Send(metric.Metric) error
}
