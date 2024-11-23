package service

import (
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

func NewStorageService(
	gaugeStorage repository.MetricRepository[model.Gauge],
	counterStorage repository.MetricRepository[model.Counter]) *StorageService {

	return &StorageService{
		gaugeStorage:   gaugeStorage,
		counterStorage: counterStorage,
	}
}

func (s StorageService) Save(metric rawMetric) error {
	if err := s.checkInput(metric); err != nil {
		return err
	}
	return s.save(metric)
}

func (s StorageService) save(metric rawMetric) error {
	saver := s.getSaverByType(metric)
	return saver.save()
}

func (s StorageService) checkInput(metric rawMetric) error {
	if metric.name == "" {
		return ErrEmptyName
	}

	if metric.kind == "" {
		return ErrEmptyKind
	}

	if metric.value == "" {
		return ErrEmptyValue
	}
	return nil
}

func (s StorageService) getSaverByType(metric rawMetric) metricSaver {
	switch metric.kind {
	case keyGauge:
		return gaugeSaver{
			metric:     metric,
			repository: s.gaugeStorage,
		}
	case keyCounter:
		return counterSaver{
			metric:     metric,
			repository: s.counterStorage,
		}
	default:
		return unknownSaver{kind: metric.kind}
	}
}
