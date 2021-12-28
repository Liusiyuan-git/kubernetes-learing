package server

import (
	"context"
	"net/http"
	"os"
	"time"
	"week2/server/middleware"

	"github.com/gorilla/mux"
)

type Server struct {
	srv *http.Server
}

var connStates = make(map[string]http.ConnState)

func NewServer(addr string) *Server {
	srv := &Server{
		srv: &http.Server{
			Addr: addr,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/get.header", srv.getHeader).Methods("POST")
	router.HandleFunc("/get.env.version", srv.getEnvVersion).Methods("POST")
	router.HandleFunc("/health", srv.health).Methods("GET")
	srv.srv.Handler = middleware.LogRequestHandler(router)
	return srv
}

func (s *Server) ListenAndServe() (<-chan error, error) {
	var err error
	errChan := make(chan error)
	go func() {
		err = s.srv.ListenAndServe()
		errChan <- err
	}()

	select {
	case err = <-errChan:
		return nil, err
	case <-time.After(time.Second):
		return errChan, nil
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *Server) getHeader(w http.ResponseWriter, req *http.Request) {
	for key, value := range req.Header {
		w.Header().Set(key, value[0])
	}
	w.Write([]byte("get header"))
}

func (s *Server) getEnvVersion(w http.ResponseWriter, req *http.Request) {
	if value, done := os.LookupEnv("VERSION"); done {
		w.Header().Set("VERSION", value)
	}
	w.Write([]byte("get env version"))
}

func (s *Server) health(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("200"))
}
