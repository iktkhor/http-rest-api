package apiserver

type Config struct {
	BinAddr string `toml:"blind_addr"`
}

func NewConfig() *Config {
	return &Config{
		BinAddr: ":8080",
	}
}