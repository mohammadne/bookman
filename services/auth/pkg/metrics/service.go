package metrics

type Service interface {
	// Save will save a Metric type
	Save(interface{})

	// Serve will serve metrics-server
	Serve()
}
