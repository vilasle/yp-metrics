package service

import (
	"testing"

	"github.com/vilasle/yp-metrics/internal/model"
	"github.com/vilasle/yp-metrics/internal/repository"
)

func Test_gaugeSaver_save(t *testing.T) {
	type fields struct {
		metric     RawMetric
		repository repository.MetricRepository[model.Gauge]
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := gaugeSaver{
				metric:     tt.fields.metric,
				repository: tt.fields.repository,
			}
			if err := s.save(); (err != nil) != tt.wantErr {
				t.Errorf("gaugeSaver.save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_counterSaver_save(t *testing.T) {
	type fields struct {
		metric     RawMetric
		repository repository.MetricRepository[model.Counter]
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := counterSaver{
				metric:     tt.fields.metric,
				repository: tt.fields.repository,
			}
			if err := s.save(); (err != nil) != tt.wantErr {
				t.Errorf("counterSaver.save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_unknownSaver_save(t *testing.T) {
	type fields struct {
		typeName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := unknownSaver{
				kind: tt.fields.typeName,
			}
			if err := c.save(); (err != nil) != tt.wantErr {
				t.Errorf("unknownSaver.save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
