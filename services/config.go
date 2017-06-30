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

func initConfig() {
	p := os.Getenv("APP_CONF_PATH")
	if p == "" {
		log.Fatal("APP_CONF_PATH not defined")
	}
	c, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(c, &conf)
}
