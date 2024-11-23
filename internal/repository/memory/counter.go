package memory

import "github.com/vilasle/yp-metrics/internal/model"

type MetricCounterMemoryRepository[T model.Counter] struct {
	metrics map[string][]model.Counter
}

func NewMetricCounterMemoryRepository() *MetricCounterMemoryRepository[model.Counter] {
	return &MetricCounterMemoryRepository[model.Counter]{
		metrics: make(map[string][]model.Counter),
	}
}

func (m *MetricCounterMemoryRepository[T]) Save(name string, metric model.Counter) error {
	if _, ok := m.metrics[name]; !ok {
		m.metrics[name] = make([]model.Counter, 0)
	}
	m.metrics[name] = append(m.metrics[name], metric)
	
	return nil
}