package rest

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/vilasle/yp-metrics/internal/metric"
	"github.com/vilasle/yp-metrics/internal/service/server"
)

type metricHandler func(http.ResponseWriter, *http.Request) http.Handler

func (mh metricHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mh(w, r)
}

func UpdateHandler(svc *server.StorageService) http.Handler {
	return metricHandler(func(w http.ResponseWriter, r *http.Request) http.Handler {
		if methodIsNotAllowed(r.Method) {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return nil
		}

		if mediaTypeIsNotAllowed(r.Header.Values("Content-Type")) {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return nil
		}
		//TODO drop it
		fmt.Println(r.URL.Path)
		//
		
		d := cleanUselessData(strings.Split(r.URL.Path, "/"))

		err := svc.Save(
			metric.NewRawMetric(getName(d), getKind(d), getValue(d)),
		)

		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(getStatusCode(err))

		return nil
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

func methodIsNotAllowed(method string) bool {
	return !methodIsAllow(method)
}

func methodIsAllow(method string) bool {
	return method == http.MethodPost
}

func mediaTypeIsNotAllowed(contentType []string) bool {
	return !mediaTypeIsAllowed(contentType)
}

func mediaTypeIsAllowed(contentType []string) bool {
	for _, ct := range contentType {
		if ct == "text/plain" {
			return true
		}
	}

	return false
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
