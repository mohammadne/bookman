package metrics

type CounterMetric interface{}

type GaugeMetric interface{}

type HistogramMetric interface {
	Finish()
}

type SummeryMetric interface {
	Finish()
}
