package tracer

import "github.com/opentracing/opentracing-go"

type Span interface {
	// SetTag sets a span to the current tag
	SetTag(string, interface{})

	// Finish will finish span scope
	Finish()

	// tracer provides access to the Tracer that created this Span.
	tracer() opentracing.Tracer

	// context provides current span context
	context() opentracing.SpanContext
}

type span struct {
	jaegerSpan opentracing.Span
}

func (span *span) SetTag(key string, value interface{}) {
	span.jaegerSpan = span.jaegerSpan.SetTag(key, value)
}

func (span *span) Finish() {
	span.jaegerSpan.Finish()
}

func (span *span) tracer() opentracing.Tracer {
	return span.jaegerSpan.Tracer()
}

func (span *span) context() opentracing.SpanContext {
	return span.jaegerSpan.Context()
}
