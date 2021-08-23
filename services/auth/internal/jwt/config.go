package jwt

type Config struct {
	AccessSecret   string `split_words:"true" required:"true"`
	AccessExpires  int    `split_words:"true" required:"true"`
	RefreshSecret  string `split_words:"true" required:"true"`
	RefreshExpires int    `split_words:"true" required:"true"`
}
