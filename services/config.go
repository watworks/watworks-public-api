package services

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

var conf *AppConfig

type AppConfig struct {
	CounterServiceUrl string `yaml:"counterServiceUrl"`
}

func Config() *AppConfig {
	if conf == nil {
		initConfig()
	}

	return conf
}

// if config file was specified, try to load it,
// otherwise, try and create config from env vars
func initConfig() {
	p := os.Getenv("APP_CONF_PATH")

	// load config from file
	if p != "" {
		c, err := ioutil.ReadFile(p)
		if err != nil {
			log.Fatal(err)
		}
		if err := yaml.Unmarshal(c, &conf); err != nil {
			log.Fatal(err)
		}
		log.Printf("\nConfig read from %s\n", p)
		return
	}

	// create config from env vars since a config file was not specified

	// ensure required keys are present
	required := []string{
		"COUNTER_SERVICE_URL",
	}
	for _, key := range required {
		if v := os.Getenv(key); v == "" {
			log.Fatalf("\n%s was not set\n", key)
		}
	}

	// create the config
	conf = &AppConfig{}
	if v := os.Getenv("COUNTER_SERVICE_URL"); v != "" {
		conf.CounterServiceUrl = v
	}
}
