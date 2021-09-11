package tracer

import (
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func New(cfg *Config) (trace.Tracer, error) {
	var exporter sdktrace.SpanExporter

	var err error
	if !cfg.Enabled {
		exporter, err = stdout.New(
			stdout.WithPrettyPrint(),
		)
	} else {
		exporter, err = jaeger.New(
			jaeger.WithAgentEndpoint(jaeger.WithAgentHost(cfg.Host), jaeger.WithAgentPort(cfg.Port)),
		)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to initialize export pipeline: %v", err)
	}

	res, err := resource.Merge(
		resource.Default(),
		resource.NewSchemaless(
			attribute.String("Namespace", cfg.Namespace),
			attribute.String("Subsystem", cfg.Subsystem),
		),
	)
	if err != nil {
		return nil, err
	}

	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(bsp),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(cfg.SampleRate))),
	)

	otel.SetTracerProvider(tp)
	var tc propagation.TraceContext
	otel.SetTextMapPropagator(tc)

	return otel.Tracer("dispatching/darius"), nil
}
