package service

import (
	"errors"
	"fmt"

	"github.com/vilasle/yp-metrics/internal/model"
	"github.com/vilasle/yp-metrics/internal/repository"
)


type metricSaver interface {
	save() error
}

type rawMetric struct {
	name  string
	kind  string
	value string
}

type gaugeSaver struct {
	metric     rawMetric
	repository repository.MetricRepository[model.Gauge]
}

func (s gaugeSaver) save() error {
	if value, err := model.GaugeFromString(s.metric.value); err == nil {
		return s.repository.Save(s.metric.name, value)
	} else {
		return err
	}
}

type counterSaver struct {
	metric     rawMetric
	repository repository.MetricRepository[model.Counter]
}

func (s counterSaver) save() error {
	if value, err := model.CounterFromString(s.metric.value); err == nil {
		return s.repository.Save(s.metric.name, value)
	} else {
		return err
	}
}

type unknownSaver struct {
	kind string
}

func (c unknownSaver) save() error {
	return errors.Join(ErrUnknownKind, fmt.Errorf("unknown type %s", c.kind))
}
