package metric

type VectorType uint8

const (
	CounterVector VectorType = iota
	HistogramVector
)

type Vector interface {
	GetLabels() []string

	// Finish will finish vector
	Finish()
}
