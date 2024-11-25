package rest

import (
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	"github.com/vilasle/yp-metrics/internal/metric"
)

type HTTPSender struct {
	*url.URL
	client http.Client
}

func NewHTTPSender(addr string) (HTTPSender, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return HTTPSender{}, err
	}
	return HTTPSender{URL: u, client: http.Client{Timeout: time.Second * 5}}, nil
}

// TODO add handling status code
func (s HTTPSender) Send(value metric.Metric) error {
	u := *s.URL
	u.Path = filepath.Join(s.Path, value.Type(), value.Name(), value.Value())

	req, err := http.NewRequest(http.MethodPost, u.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "text/plain")
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
