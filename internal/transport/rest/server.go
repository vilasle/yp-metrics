package rest

import (
	"context"
	"net/http"
	"time"
)

type HTTPServer struct {
	srv     *http.Server
	mux     *http.ServeMux
	running bool
}

func NewHTTPServer(addr string) HTTPServer {
	return HTTPServer{
		srv: &http.Server{
			Addr:         addr,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
		mux: http.NewServeMux(),
	}
}

func (s HTTPServer) Register(path string, handler http.Handler) {
	s.mux.Handle(path, handler)
}

func (s *HTTPServer) Start() error {
	s.srv.Handler = s.mux
	s.running = true
	defer func() {
		s.running = false
	}()

	err := s.srv.ListenAndServe()

	if err != nil && err == http.ErrServerClosed {
		err = nil
	}

	return err
}

func (s HTTPServer) IsRunning() bool {
	return s.running
}

func (s *HTTPServer) Stop() error {
	if err := s.srv.Shutdown(context.Background()); err != nil {
		s.running = false
		return err
	}
	return nil
}

func (s HTTPServer) ForceStop() error {
	return s.srv.Close()
}