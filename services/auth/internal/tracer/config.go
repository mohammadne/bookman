package tracer

type Config struct {
	Enabled      bool    `split_words:"false" default:"true"`
	ServiceName  string  `default:"soteria" split_words:"true"`
	SamplerType  string  `default:"const" split_words:"true"`
	SamplerParam float64 `default:"1" split_words:"true"`
	Host         string  `default:"localhost" split_words:"true"`
	Port         int     `default:"6831" split_words:"true"`
}
