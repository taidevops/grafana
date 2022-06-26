package tracing

import "github.com/taidevops/grafana/pkg/setting"

const (
	envJeagerAgentHost = "JAEGER_AGENT_HOST"
	envJaegerAgentPort = "JAEGER_AGENT_PORT"
)

func ProvideService(cfg *setting.Cfg) (Tracer, error) {
	ts, ots, err :=
}

func parseSettings(cfg *setting.Cfg) (*OpenTracing, *Opentelemetry, error) {
	
}
