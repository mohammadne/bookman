package database

type Config struct {
	Username string
	Password string
	Host     string `default:"127.0.0.1:3306"`
	Schema   string `default:"bookman"`
}
