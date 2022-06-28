package telegramserver

type Config struct {
	Token         string `toml:"token"`
	ResourcesPath string `toml:"resources_path"`
}

func NewConfig() *Config {
	return &Config{
		Token:         "",
		ResourcesPath: "./resources/",
	}
}
