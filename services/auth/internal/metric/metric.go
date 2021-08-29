package metric

import "github.com/mohammadne/go-pkgs/logger"

type Metric interface {
	StartVector(VectorType, string) Vector
}

type prometheusMetric struct {
	config *Config
	logger *logger.Logger
}

var singleton *prometheusMetric

func NewPrometheus(cfg *Config, logger *logger.Logger) Metric {
	if singleton == nil {
		singleton = &prometheusMetric{config: cfg, logger: logger}
	} else {
		return singleton
	}

	return singleton
}

func (prom *prometheusMetric) StartVector(vectorType VectorType, module string) Vector {
	return nil
}
