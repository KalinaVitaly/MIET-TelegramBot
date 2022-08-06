package telegramserver

type Config struct {
	Token         string `toml:"token"`
	ResourcesPath string `toml:"resources_path"`
	Host          string `toml:"host"`
	Port          string `toml:"port"`
	User          string `toml:"user"`
	Password      string `toml:"password"`
	NameDB        string `toml:"dbname"`
}

func NewConfig() *Config {
	return &Config{
		Token:         "",
		ResourcesPath: "./resources/",
		Host:          "",
		Port:          "",
		User:          "",
		Password:      "",
		NameDB:        "",
	}
}
