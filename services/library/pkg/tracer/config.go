package tracer

type Config struct {
	Enabled    bool    `default:"false"`
	Host       string  `split_words:"true"`
	Port       string  `split_words:"true"`
	SampleRate float64 `split_words:"true" default:"0.1"`
	Namespace  string  `required:"true"`
	Subsystem  string  `required:"true"`
}
