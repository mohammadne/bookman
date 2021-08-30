package tracer

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

type Tracer interface {
	StartSpan(string) Span
	StartSpanFromRootSpan(Span, string) Span
	Close() error
}

type jaeger struct {
	tracer opentracing.Tracer
	closer io.Closer
}

func NewJaeger(cfg *Config) (Tracer, error) {
	tracer, closer, err := config.Configuration{
		ServiceName: cfg.ServiceName,
		Disabled:    !cfg.Enabled,
		Sampler: &config.SamplerConfig{
			Type:  cfg.SamplerType,
			Param: cfg.SamplerParam,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		},
	}.NewTracer()

	if err != nil {
		return nil, fmt.Errorf("failed to create new tracer: %w", err)
	}

	return &jaeger{tracer: tracer, closer: closer}, nil
}

func (j *jaeger) StartSpan(operation string) Span {
	return &span{
		jaegerSpan: j.tracer.StartSpan(operation),
	}
}

func (j *jaeger) StartSpanFromRootSpan(rootSpan Span, operation string) Span {
	return &span{
		jaegerSpan: rootSpan.tracer().StartSpan(
			operation,
			opentracing.ChildOf(rootSpan.context()),
		),
	}
}

func (j *jaeger) Close() error {
	return j.closer.Close()
}
