package metric

import (
	"time"

	"github.com/mohammadne/bookman/auth/internal"
	"github.com/mohammadne/go-pkgs/logger"
	"github.com/prometheus/client_golang/prometheus"
)

type Metric interface {
	StartCounterVector(string, string) *statusVector
	StartHistogramVector(string, string) *timeVector
}

type prometheusMetric struct {
	config *Config
	logger *logger.Logger

	counterVec   *prometheus.CounterVec
	histogramVec *prometheus.HistogramVec
}

var singleton *prometheusMetric

func NewPrometheus(cfg *Config, logger *logger.Logger) Metric {
	if singleton == nil {
		singleton = &prometheusMetric{config: cfg, logger: logger}
	} else {
		return singleton
	}

	singleton.counterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: internal.Namespace,
			Subsystem: internal.Subsystem,
			Name:      "counterVec",
		},
		statusVector{}.getLabels(),
	)

	singleton.histogramVec = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: internal.Namespace,
			Subsystem: internal.Subsystem,
			Name:      "histogramVec",
		},
		timeVector{}.getLabels(),
	)

	prometheus.MustRegister(singleton.counterVec)
	prometheus.MustRegister(singleton.histogramVec)

	return singleton
}

func (prom *prometheusMetric) StartCounterVector(module string, function string) *statusVector {
	return &statusVector{
		vector:   prom.counterVec,
		module:   module,
		function: function,
	}
}

func (prom *prometheusMetric) StartHistogramVector(module string, function string) *timeVector {
	return &timeVector{
		vector:   prom.histogramVec,
		start:    time.Now(),
		module:   module,
		function: function,
	}
}
