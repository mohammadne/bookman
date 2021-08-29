package metrics

import (
	"errors"
	"net/http"

	"github.com/mohammadne/bookman/auth/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.uber.org/zap"
)

type prometheusMetric struct {
	config *Config
	logger *zap.Logger

	// prometheus vectors
	// call        *prometheus.SummaryVec
	// cmq         *prometheus.CounterVec
	// emq         *prometheus.SummaryVec
	// handlers    *prometheus.SummaryVec
	// janus       *prometheus.SummaryVec
	// middlewares *prometheus.SummaryVec
	// ripo        *prometheus.SummaryVec
}

var singleton *prometheusMetric

// NewPrometheus is a singleton factory function responsible
func NewPrometheus(cfg *Config, log *zap.Logger) metrics.Service {
	if singleton != nil {
		return singleton
	}

	singleton = &prometheusMetric{config: cfg, logger: log}

	// singleton.call = prometheus.NewSummaryVec(
	// 	singleton.summaryOpts("call"),
	// 	[]string{"reciever", "method"},
	// )

	// singleton.cmq = prometheus.NewCounterVec(
	// 	singleton.counterOpts("cmq"),
	// 	[]string{"subject"},
	// )

	// singleton.emq = prometheus.NewSummaryVec(
	// 	singleton.summaryOpts("emq"),
	// 	[]string{"reciever", "method"},
	// )

	// singleton.handlers = prometheus.NewSummaryVec(
	// 	singleton.summaryOpts("handlers"),
	// 	[]string{"reciever", "method", "status_code"},
	// )

	// singleton.janus = prometheus.NewSummaryVec(
	// 	singleton.summaryOpts("janus"),
	// 	[]string{"reciever", "method"},
	// )

	// singleton.middlewares = prometheus.NewSummaryVec(
	// 	singleton.summaryOpts("middlewares"),
	// 	[]string{"function", "status_code"},
	// )

	// singleton.ripo = prometheus.NewSummaryVec(
	// 	singleton.summaryOpts("ripo"),
	// 	[]string{"function"},
	// )

	// prometheus.MustRegister(singleton.call)
	// prometheus.MustRegister(singleton.cmq)
	// prometheus.MustRegister(singleton.emq)
	// prometheus.MustRegister(singleton.handlers)
	// prometheus.MustRegister(singleton.janus)
	// prometheus.MustRegister(singleton.middlewares)
	// prometheus.MustRegister(singleton.ripo)

	return singleton
}

func (p *prometheusMetric) counterOpts(name string) prometheus.CounterOpts {
	return prometheus.CounterOpts{
		Namespace: p.config.Namespace,
		Subsystem: p.config.Subsystem,
		Name:      name,
	}
}

func (p *prometheusMetric) summaryOpts(name string) prometheus.SummaryOpts {
	return prometheus.SummaryOpts{
		Namespace: p.config.Namespace,
		Subsystem: p.config.Subsystem,
		Name:      name,
	}
}

func (p *prometheusMetric) Save(m interface{}) {
	// 	switch metric := m.(type) {
	// 	case call.Metric:
	// 		p.call.
	// 			WithLabelValues(metric.Reciever, metric.Method).
	// 			Observe(metric.Timing.Duration())
	// 	case cmq.Metric:
	// 		p.cmq.
	// 			WithLabelValues(metric.Subject).
	// 			Inc()
	// 	case emq.Metric:
	// 		p.call.
	// 			WithLabelValues(metric.Reciever, metric.Method).
	// 			Observe(metric.Timing.Duration())
	// 	case handlers.Metric:
	// 		p.call.
	// 			WithLabelValues(metric.Reciever, metric.Method, metric.StatusCode).
	// 			Observe(metric.Timing.Duration())
	// 	case janus.Metric:
	// 		p.call.
	// 			WithLabelValues(metric.Reciever, metric.Method).
	// 			Observe(metric.Timing.Duration())
	// 	case middlewares.Metric:
	// 		p.call.
	// 			WithLabelValues(metric.Function, metric.StatusCode).
	// 			Observe(metric.Timing.Duration())
	// 	case ripo.Metric:
	// 		p.call.
	// 			WithLabelValues(metric.Function).
	// 			Observe(metric.Timing.Duration())
	// 	}
}

func (p *prometheusMetric) Serve() {
	srv := new(http.ServeMux)
	srv.Handle("/metrics", promhttp.Handler())

	go func() {
		err := http.ListenAndServe(p.config.Address, srv)
		if !errors.Is(err, http.ErrServerClosed) {
			p.logger.Fatal("metric server initiation failed", zap.Error(err))
		}
	}()
}
