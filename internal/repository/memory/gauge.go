package memory

import "github.com/vilasle/yp-metrics/internal/model"

type MetricGaugeMemoryRepository[T model.Gauge] struct {
	metrics map[string]T
}

func NewMetricGaugeMemoryRepository() *MetricGaugeMemoryRepository[model.Gauge] {
	return &MetricGaugeMemoryRepository[model.Gauge]{
		metrics: make(map[string]model.Gauge),
	}
}

func (m *MetricGaugeMemoryRepository[T]) Save(name string, metric T) error{
	m.metrics[name] = metric
	return nil
}