package telegramserver

type Config struct {
	Token string `toml:"token"`
}

func NewConfig() *Config {
	return &Config{
		Token: "",
	}
}
