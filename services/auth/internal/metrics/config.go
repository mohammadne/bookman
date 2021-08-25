package metrics

type Config struct {
	Address   string `default:"localhost:9090"`
	Namespace string `default:"dispatching"`
	Subsystem string `default:"bell"`
}
