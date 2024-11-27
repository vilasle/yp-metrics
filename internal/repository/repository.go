package repository

import (
	"github.com/vilasle/yp-metrics/internal/model"
)

type MetricRepository[T model.Gauge | model.Counter] interface {
	Save(string, T) error
	Get(string) (T, error)
	All() (map[string]T, error)
}