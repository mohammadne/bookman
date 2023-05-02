package database

type Config struct {
	Driver       string `default:"postgres"`
	Host         string `default:"127.0.0.1"`
	Port         string `default:"5432"`
	User         string `default:"admin"`
	Password     string `default:"admin"`
	DatabaseName string `default:"bell" split_words:"true"`
	SSLMode      string `default:"disable" split_words:"true"`
}
