package api

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"

	"github.com/grafana/grafana/pkg/web"
)

type HTTPServer struct {
	web *web.Mux
	context context.Context
	
}