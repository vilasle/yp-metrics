package service

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vilasle/yp-metrics/internal/model"
	"github.com/vilasle/yp-metrics/internal/repository"
	"github.com/vilasle/yp-metrics/internal/repository/memory"
)

func TestStorageService_Save(t *testing.T) {
	type fields struct {
		gaugeStorage   repository.MetricRepository[model.Gauge]
		counterStorage repository.MetricRepository[model.Counter]
	}

	_fields := fields{
		gaugeStorage:   memory.NewMetricGaugeMemoryRepository(),
		counterStorage: memory.NewMetricCounterMemoryRepository(),
	}

	type args struct {
		metric RawMetric
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		err    error
	}{
		{
			name:   "problems with input, did not fill name of value",
			fields: _fields,
			args: args{
				metric: RawMetric{
					name:  "",
					kind:  keyGauge,
					value: "12.45",
				},
			},
			err: ErrEmptyName,
		},
		{
			name:   "problems with input, did not fill kind of metric",
			fields: _fields,
			args: args{
				metric: RawMetric{
					name:  "test",
					kind:  "",
					value: "12.45",
				},
			},
			err: ErrEmptyKind,
		},
		{
			name:   "problems with input, did not fill value of metric",
			fields: _fields,
			args: args{
				metric: RawMetric{
					name:  "test",
					kind:  keyGauge,
					value: "",
				},
			},
			err: ErrEmptyValue,
		},
		{
			name:   "metric is filled but contents unknown kind",
			fields: _fields,
			args: args{
				metric: RawMetric{
					name:  "test",
					kind:  "test",
					value: "123.4",
				},
			},
			err: ErrUnknownKind,
		},
		{
			name:   "gauge metric is filled and kind is right",
			fields: _fields,
			args: args{
				metric: RawMetric{
					name:  "test",
					kind:  keyGauge,
					value: "123.4",
				},
			},
			err: nil,
		},
		{
			name:   "gauge metric is filled and kind is right, value has wrong type",
			fields: _fields,
			args: args{
				metric: RawMetric{
					name:  "test",
					kind:  keyGauge,
					value: "test",
				},
			},
			err: model.ErrConvertMetricFromString,
		},
		{
			name:   "counter metric is filled and kind is right",
			fields: _fields,
			args: args{
				metric: RawMetric{
					name:  "test",
					kind:  keyCounter,
					value: "144",
				},
			},
			err: nil,
		},
		{
			name:   "counter metric is filled and kind is right, value has wrong type",
			fields: _fields,
			args: args{
				metric: RawMetric{
					name:  "test",
					kind:  keyCounter,
					value: "test",
				},
			},
			err: model.ErrConvertMetricFromString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStorageService(tt.fields.gaugeStorage, tt.fields.counterStorage)
			err := s.Save(tt.args.metric)
			if tt.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorContains(t, s.Save(tt.args.metric), tt.err.Error())
			}
		})
	}
}

func TestStorageService_checkInput(t *testing.T) {
	type fields struct {
		gaugeStorage   repository.MetricRepository[model.Gauge]
		counterStorage repository.MetricRepository[model.Counter]
	}

	_fields := fields{
		gaugeStorage:   memory.NewMetricGaugeMemoryRepository(),
		counterStorage: memory.NewMetricCounterMemoryRepository(),
	}

	type args struct {
		metric RawMetric
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		err    error
	}{
		{
			name:   "all is filled",
			fields: _fields,
			args: args{
				metric: RawMetric{
					name:  "test",
					kind:  "gauge",
					value: "12.45",
				},
			},
			err: nil,
		},
		{
			name:   "empty name",
			fields: _fields,
			args: args{
				metric: RawMetric{
					name:  "",
					kind:  "gauge",
					value: "12.45",
				},
			},
			err: ErrEmptyName,
		},
		{
			name:   "empty kind",
			fields: _fields,
			args: args{
				metric: RawMetric{
					name:  "test",
					kind:  "",
					value: "12.45",
				},
			},
			err: ErrEmptyKind,
		},
		{
			name:   "empty value",
			fields: _fields,
			args: args{
				metric: RawMetric{
					name:  "test",
					kind:  "gauge",
					value: "",
				},
			},
			err: ErrEmptyValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StorageService{
				gaugeStorage:   tt.fields.gaugeStorage,
				counterStorage: tt.fields.counterStorage,
			}
			err := s.checkInput(tt.args.metric)

			if tt.err == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorContains(t, err, tt.err.Error())
			}
		})
	}
}

func TestStorageService_getSaverByType(t *testing.T) {
	type fields struct {
		gaugeStorage   repository.MetricRepository[model.Gauge]
		counterStorage repository.MetricRepository[model.Counter]
	}

	_fields := fields{
		gaugeStorage:   memory.NewMetricGaugeMemoryRepository(),
		counterStorage: memory.NewMetricCounterMemoryRepository(),
	}

	type args struct {
		metric RawMetric
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   metricSaver
	}{
		{
			name:   "gauge",
			fields: _fields,
			args: args{
				metric: RawMetric{
					name:  "test",
					kind:  "gauge",
					value: "12.45",
				},
			},
			want: gaugeSaver{
				repository: _fields.gaugeStorage,
				metric: RawMetric{
					name:  "test",
					kind:  "gauge",
					value: "12.45",
				},
			},
		},
		{
			name:   "counter",
			fields: _fields,
			args: args{
				metric: RawMetric{
					name:  "test",
					kind:  "counter",
					value: "12",
				},
			},
			want: counterSaver{
				repository: _fields.counterStorage,
				metric: RawMetric{
					name:  "test",
					kind:  "counter",
					value: "12",
				},
			},
		},
		{
			name:   "unknown saver",
			fields: _fields,
			args: args{
				metric: RawMetric{
					name:  "test",
					kind:  "test",
					value: "12",
				},
			},
			want: unknownSaver{
				kind: "test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StorageService{
				gaugeStorage:   tt.fields.gaugeStorage,
				counterStorage: tt.fields.counterStorage,
			}
			got := s.getSaverByType(tt.args.metric)

			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("StorageService.getSaverByType() = %v, want %v", got, tt.want)
			}
		})
	}
}
