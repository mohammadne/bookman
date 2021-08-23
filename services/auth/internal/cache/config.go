package cache

import "time"

// Config is the config of the cache client.
type Config struct {
	Mode               int           `split_words:"true" default:"0"`
	URL                string        `split_words:"true"`
	MasterURL          string        `split_words:"true"`
	SlaveURL           string        `split_words:"true"`
	Password           string        `default:"" split_words:"true"`
	ExpirationTime     int           `default:"10" split_words:"true"`
	PoolSize           int           `split_words:"true" default:"10"`
	MaxRetries         int           `split_words:"true" default:"0"`
	ReadTimeout        time.Duration `split_words:"true" default:"3s"`
	PoolTimeout        time.Duration `split_words:"true" default:"4s"`
	MinRetryBackoff    time.Duration `split_words:"true" default:"8ms"`
	MaxRetryBackoff    time.Duration `split_words:"true" default:"512ms"`
	IdleTimeout        time.Duration `split_words:"true" default:"300s"`
	IdleCheckFrequency time.Duration `split_words:"true" default:"60s"`
	SetMemberExpTime   time.Duration `split_words:"true" default:"300s"`
}
