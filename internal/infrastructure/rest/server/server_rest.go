package server

import (
	"context"
	"net/http"
	"time"

	"github.com/ninja-dark/fibonacci_testtask/internal/cases"
)

type ServerRest struct {
	srv http.Server
	fibo *cases.Fibo
}

func NewServer (addr string, h http.Handler) *ServerRest {
	s := &ServerRest{}

	s.srv = http.Server{
		Addr:              addr,
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
	return s
}

func (s *NewServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	s.srv.Shutdown(ctx)
	cancel()
}

func (s *NewServer) Start(fibo *cases.Fibo) {
	s.us = fibo
	
	go s.srv.ListenAndServe()
}