package routing

import "github.com/taidevops/grafana/pkg/web"

type Router interface {
	Handle(method, pattern string, handlers []web.Handler)
	Get(pattern string, handlers ...web.Handler)
}

type RouteRegister interface {
	Get(string, ...web.Handler)

	Post(string, ...web.Handler)

	Delete(string, ...web.Handler)
}
