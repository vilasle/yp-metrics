package model

import (
	"strconv"

	"errors"
)

// may be need to change it as struct
type Gauge float64

func GaugeFromString(s string) (Gauge, error) {
	if v, err := strconv.ParseFloat(s, 64); err != nil {
		return Gauge(0), errors.Join(err, ErrConvertMetricFromString)
	} else {
		return Gauge(v), nil
	}
}

type Counter int64

func CounterFromString(s string) (Counter, error) {
	if v, err := strconv.ParseInt(s, 10, 64); err != nil {
		return Counter(0), errors.Join(err, ErrConvertMetricFromString)
	} else {
		return Counter(v), nil
	}
}
