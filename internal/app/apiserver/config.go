package apiserver

type Config struct {
	//Адрес на котором запускаем веб-сервер
	BindAddr string `toml:"bind_addr"`
}


func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}