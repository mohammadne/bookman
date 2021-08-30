package metric

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type VectorType uint8

const (
	CounterVector VectorType = iota
	HistogramVector
)

type Vector interface {
	// Finish will finish vector
	Finish()

	getLabels() []string
}

// ===========================================================> counterVector

type counterVector struct {
	vector *prometheus.CounterVec

	// labels
	module   string
	function string
}

func (v *counterVector) Finish() {
	v.vector.
		WithLabelValues(v.module, v.function).
		Inc()
}

func (v counterVector) getLabels() []string {
	return []string{"module", "function"}
}

// ===========================================================> histogramVector

type histogramVector struct {
	vector *prometheus.HistogramVec
	start  time.Time

	// labels
	module   string
	function string
}

func (v *histogramVector) Finish() {
	v.vector.
		WithLabelValues(v.module, v.function).
		Observe(time.Since(v.start).Seconds())
}

func (v histogramVector) getLabels() []string {
	return []string{"module", "function"}
}
