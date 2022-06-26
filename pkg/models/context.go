package models

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/taidevops/grafana/pkg/infra/log"
	"github.com/taidevops/grafana/pkg/setting"
	"github.com/taidevops/grafana/pkg/web"
)

type ReqContext struct {
	*web.Context
	*SignedInUser
	UserToken *UserToken

	IsSignedIn     bool
	IsRenderCall   bool
	AllowAnonymous bool
	SkipCache      bool
	Logger         log.Logger
	// RequestNonce is a cryptographic request identifier for use with Content Security Policy.
	RequestNonce          string
	IsPublicDashboardView bool

	PerfmonTimer   prometheus.Summary
	LookupTokenErr error
}

func (ctx *ReqContext) Handle(cfg *setting.Cfg, status int, title string, err error) {
	data := struct {
		Title     string
		AppTitle  string
		AppSubUrl string
		Theme     string
		ErrorMsg  error
	}{title, "Grafana", cfg.AppSubURL, "dark", nil}
	if err != nil {
		if setting.Env != setting.Prod {
			data.ErrorMsg = err
		}
	}

	ctx.HTML(status, cfg.ErrTemplateName, data)
}

func (ctx *ReqContext) IsApiRequest() bool {
	return strings.HasPrefix(ctx.Req.URL.Path, "/api")
}

func (ctx *ReqContext) JsonApiErr(status int, message string, err error) {
	resp := make(map[string]interface{})

}
