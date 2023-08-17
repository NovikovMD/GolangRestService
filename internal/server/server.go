package server

import (
	"context"
	"net/http"
	"time"
)

type HttpServer struct {
	Server *http.Server
}

func InitServer(port string, handler http.Handler) *HttpServer {
	return &HttpServer{
		Server: &http.Server{
			Addr:         ":" + port,
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  10 * time.Second,
			Handler:      handler,
		},
	}
}

func (hs *HttpServer) Run() error {
	return hs.Server.ListenAndServe()
}

func (hs *HttpServer) Shutdown(context context.Context) error {
	return hs.Server.Shutdown(context)
}
