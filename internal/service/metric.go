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

type RawMetric struct {
	name  string
	kind  string
	value string
}

func NewRawMetric(kind, name, value string) RawMetric {
	return RawMetric{name: name, kind: kind, value: value}
}

type gaugeSaver struct {
	metric     RawMetric
	repository repository.MetricRepository[model.Gauge]
}

func (s gaugeSaver) save() error {
	if value, err := model.GaugeFromString(s.metric.value); err == nil {
		return s.repository.Save(s.metric.name, value)
	} else {
		return errors.Join(err, ErrInvalidValue)
	}
}

type counterSaver struct {
	metric     RawMetric
	repository repository.MetricRepository[model.Counter]
}

func (s counterSaver) save() error {
	if value, err := model.CounterFromString(s.metric.value); err == nil {
		return s.repository.Save(s.metric.name, value)
	} else {
		return errors.Join(err, ErrInvalidValue)
	}
}

type unknownSaver struct {
	kind string
}

func (c unknownSaver) save() error {
	return errors.Join(ErrUnknownKind, fmt.Errorf("unknown type %s", c.kind))
}
