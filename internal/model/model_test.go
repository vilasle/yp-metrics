package model

import (
	"testing"
)

func TestGaugeFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Gauge
		wantErr bool
	}{
		{
			name: "valid gauge",
			args: args{
				s: "1234567890.545342",
			},
			want:	Gauge(1234567890.545342),
			wantErr: false,
		},
		{
			name: "invalid gauge",
			args: args{
				s: "1234567890a",
			},
			want:	Gauge(0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GaugeFromString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("GaugeFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GaugeFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCounterFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Counter
		wantErr bool
	}{
		{
			name: "valid counter",
			args: args{ "1465645" },
			want: Counter(1465645),
			wantErr: false,
		},
		{
			name: "invalid counter",
			args: args{ "14gd65g64d5g" },
			want: Counter(0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CounterFromString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("CounterFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CounterFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
