package server

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"sync"


	"github.com/grafana/grafana/pkg/api"
	"github.com/grafana/grafana/pkg/setting"
	"golang.org/x/sync/errgroup"
)

type Options struct {

}

func New(opts Options, cfg * setting.Cfg, httpServer *api.HTTPServer) (*Server, error) {
	s, err := newServer
}

func newServer(opts Options, cfg *setting.Cfg, httpServer *api.HTTPServer) (*Server, error) {
	rootCtx, shutdownFn := context.WithCancel(context.Background())
	childRoutines, childCtx := errgroup.WithContext(rootCtx)

	s := &Server{
		context: childCtx
	}

	return s, nil
}

// Server is responsible for managing the lifecycle of services.
type Server struct {
	context          context.Context
	shutdownFn       context.CancelFunc

	isInitialized    bool

	HTTPServer *api.HTTPServer
}

func (s *Server) init() error {
	if s.isInitialized {
		return nil
	}
	s.isInitialized = true


}

// writePIDFile retrieves the current process ID and writes it to file.
func (s *Server) writePIDFile() {
	if s.pidFile == "" {
		return
	}

	// Ensure the required directory structure exists.
	err := os.MkdirAll(filepath.Dir(s.pidFile), 0700)
	if err != nil {
		s.log.Error("Failed to verify pid directory", "error", err)
		os.Exit(1)
	}

	// Retrieve the PID and write it to file.
	pid := strconv.Itoa(os.Getpid())
	if err := ioutil.WriteFile(s.pidFile, []byte(pid), 0644); err != nil {
		s.log.Error("Failed to write pidfile", "error", err)
		os.Exit(1)
	}

	s.log.Info("Writing PID file", "path", s.pidFile, "pid", pid)
}
