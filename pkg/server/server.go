package server

import (
	"context"
	"fmt"
	"sync"
)

func New() (*Server, error) {

}

type Server struct {
	context          context.Context
	isInitialized    bool
	mtx              sync.Mutex
}

// init initializes the server and its services.
func (s *Server) init() error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if s.isInitialized {
		return nil
	}
	s.isInitialized = true


}

func (s *Server) Run() error {

}
