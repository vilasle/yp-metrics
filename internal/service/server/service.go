package server

import (
	"errors"
	"fmt"

	"github.com/vilasle/yp-metrics/internal/metric"
	"github.com/vilasle/yp-metrics/internal/model"
	"github.com/vilasle/yp-metrics/internal/repository"
)

const (
	keyGauge   = "gauge"
	keyCounter = "counter"
)

type StorageService struct {
	gaugeStorage   repository.MetricRepository[model.Gauge]
	counterStorage repository.MetricRepository[model.Counter]
}

type metricSaver interface {
	Save() error
}

type unknownSaver struct {
	kind string
}

func (c unknownSaver) Save() error {
	return errors.Join(ErrUnknownKind, fmt.Errorf("unknown type %s", c.kind))
}

func NewStorageService(
	gaugeStorage repository.MetricRepository[model.Gauge],
	counterStorage repository.MetricRepository[model.Counter]) *StorageService {

	return &StorageService{
		gaugeStorage:   gaugeStorage,
		counterStorage: counterStorage,
	}
}

func (s StorageService) Save(data metric.RawMetric) error {
	if err := s.checkInput(data); err != nil {
		return err
	}
	return s.save(data)
}

func (s StorageService) save(data metric.RawMetric) error {
	saver := s.getSaverByType(data)

	err := saver.Save()
	if errors.Is(err, metric.ErrConvertingRawValue) {
		return errors.Join(err, ErrInvalidValue)
	} else if errors.Is(err, metric.ErrUnknownKind) {
		return errors.Join(err, ErrUnknownKind)
	}

	return err
}

func (s StorageService) checkInput(data metric.RawMetric) error {
	if data.Kind == "" {
		return ErrEmptyKind
	}

	if data.Name == "" {
		return ErrEmptyName
	}

	if data.Value == "" {
		return ErrEmptyValue
	}
	return nil
}

func (s StorageService) getSaverByType(data metric.RawMetric) metricSaver {
	switch data.Kind {
	case keyGauge:
		return metric.NewGaugeSaver(data, s.gaugeStorage)
	case keyCounter:
		return metric.NewCounterSaver(data, s.counterStorage)
	default:
		return unknownSaver{kind: data.Kind}
	}
}
