package rest

import (
	"errors"
	"net/http"
	"strings"

	"github.com/vilasle/yp-metrics/internal/metric"
	"github.com/vilasle/yp-metrics/internal/service/server"
)

func UpdateHandler(svc *server.StorageService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d := cleanUselessData(strings.Split(r.URL.Path, "/"))

		err := svc.Save(
			metric.NewRawMetric(getName(d), getKind(d), getValue(d)),
		)

		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(getStatusCode(err))
	})
}

func DisplayAllMetrics(svc *server.StorageService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotImplemented)
	})
}

func DisplayMetric(svc *server.StorageService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain]; charset=utf-8")
		w.WriteHeader(http.StatusNotImplemented)
	})
}

func cleanUselessData(data []string) []string {
	startInx := 0
	for startInx = range data {
		if data[startInx] == "" || data[startInx] == "update" {
			continue
		}
		return data[startInx:]
	}
	return []string{}
}

func getKind(data []string) string {
	if len(data) > 0 {
		return data[0]
	}
	return ""
}

func getName(data []string) string {
	if len(data) > 1 {
		return data[1]
	}
	return ""
}

func getValue(data []string) string {
	if len(data) > 2 {
		return data[2]
	}
	return ""
}

func getStatusCode(err error) int {
	if errorBadRequest(err) {
		return http.StatusBadRequest
	} else if errorNotFound(err) {
		return http.StatusNotFound
	}
	return http.StatusOK
}

func errorBadRequest(err error) bool {
	return errors.Is(err, server.ErrEmptyKind) ||
		errors.Is(err, server.ErrUnknownKind) ||
		errors.Is(err, server.ErrInvalidValue) ||
		errors.Is(err, server.ErrEmptyValue)
}

func errorNotFound(err error) bool {
	return err == server.ErrEmptyName
}
