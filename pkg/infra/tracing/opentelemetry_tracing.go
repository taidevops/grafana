package tracing

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	trace "go.opentelemetry.io/otel/trace"
)

const (
	jaegerExporter string = "jaeger"
	otlpExporter   string = "otlp"
	noopExporter   string = "noop"

	jaegerPropagator string = "jaeger"
	w3cPropagator    string = "w3c"
)

type Tracer interface {
	Run(context.Context) error
	Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, Span)
	Inject(context.Context, http.Header, Span)
}

type Span interface {
	End()
	SetAttributes(key string, value interface{}, kv attribute.KeyValue)
	SetName(name string)
	SetStatus(code codes.Code, description string)
	RecordError(err error, options ...trace.EventOption)
	AddEvents(keys []string, values []EventValue)
}

type EventValue struct {
	Str string
	Num int64
}
