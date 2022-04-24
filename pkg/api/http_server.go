package api

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/grafana/grafana/pkg/setting"
	"github.com/grafana/grafana/pkg/web"
)

type HTTPServer struct {
	web         *web.Mux
	context     context.Context
	httpSrv     *http.Server
	middlewares web.Handler

	Cfg *setting.Cfg
}

type ServerOptions struct {
	Listener net.Listener
}

func ProvideHTTPServer() (*HTTPServer, error) {
	web.Env = "prod"
	m := web.New()

	hs := &HTTPServer{
		web: m,
	}

	return hs, nil
}

func (hs *HTTPServer) Run(ctx context.Context) error {
	hs.context = ctx

	hs.httpSrv = &http.Server{
		Addr:        net.JoinHostPort("", ""),
		Handler:     hs.web,
		ReadTimeout: 10 * time.Second,
	}
	return nil
}

func (hs *HTTPServer) getListener() (net.Listener, error) {

}
