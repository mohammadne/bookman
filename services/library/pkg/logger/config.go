package logger

// Config is the configs needed for logging purposes.
type Config struct {
	Development      bool   `default:"true"`
	EnableCaller     bool   `default:"false" split_words:"true"`
	EnableStacktrace bool   `default:"false" split_words:"true"`
	Encoding         string `default:"console"`
	Level            string `default:"warn"`
}
