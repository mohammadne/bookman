package metric_vectors

import "github.com/prometheus/client_golang/prometheus"

type counterVector struct {
	vector   *prometheus.CounterVec
	module   string
	function string
}

func (v *counterVector) GetLabels() []string {
	return []string{"module", "function"}
}

func (v *counterVector) Finish() {
	v.vector.
		WithLabelValues(v.module, v.function).
		Inc()
}
