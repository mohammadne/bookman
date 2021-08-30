package metric

import (
	"time"

	"github.com/mohammadne/bookman/auth/internal"
	"github.com/mohammadne/go-pkgs/logger"
	"github.com/prometheus/client_golang/prometheus"
)

type Metric interface {
	StartCounterVector(string) *counterVector
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
		counterVector{}.getLabels(),
	)

	singleton.histogramVec = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: internal.Namespace,
			Subsystem: internal.Subsystem,
			Name:      "histogramVec",
		},
		histogramVector{}.getLabels(),
	)

	return singleton
}

func (prom *prometheusMetric) StartCounterVector(module string, function string) *counterVector {
	return &counterVector{
		vector:   prom.counterVec,
		module:   module,
		function: function,
	}
}

func (prom *prometheusMetric) StartHistogramVector(module string, function string) *histogramVector {
	return &histogramVector{
		vector:   prom.histogramVec,
		start:    time.Now(),
		module:   module,
		function: function,
	}
}
