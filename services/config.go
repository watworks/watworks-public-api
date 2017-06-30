package services

var conf *AppConfig

type AppConfig struct {
	ServerUrl         string
	CounterServiceUrl string
}

func Config() *AppConfig {
	if conf == nil {
		initConfig()
	}

	return conf
}

// TODO: change this to load a config file instead - expect the filepath as an env variable?
func initConfig() {
	conf = &AppConfig{":80", "http://localhost:8006/"}
}
