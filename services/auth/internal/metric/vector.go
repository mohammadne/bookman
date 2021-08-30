package metric

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type VectorType uint8

const (
	CounterVector VectorType = iota
	HistogramVector

	//

	// statusVector
	// timeVector
)

type Vector interface {
	// Collect will finish vector
	Collect()

	getLabels() []string
}

// ===========================================================> statusVector

type Status string

const (
	Success Status = "success"
	Failure Status = "failure"
)

type statusVector struct {
	vector *prometheus.CounterVec

	// labels
	module   string
	function string
	status   Status
}

func (v *statusVector) Collect() {
	v.vector.
		WithLabelValues(v.module, v.function).
		Inc()
}

func (v *statusVector) SetStatus(status Status) {
	v.status = status
}

func (v statusVector) getLabels() []string {
	return []string{"module", "function"}
}

// ===========================================================> timeVector

type timeVector struct {
	vector *prometheus.HistogramVec
	start  time.Time

	// labels
	module   string
	function string
}

func (v *timeVector) Collect() {
	v.vector.
		WithLabelValues(v.module, v.function).
		Observe(time.Since(v.start).Seconds())
}

func (v timeVector) getLabels() []string {
	return []string{"module", "function"}
}
