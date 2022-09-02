package utils

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Db struct {
		User     string `yaml:postgres`
		Password string `yaml:password`
		Port     string `yaml:port`
	} `yaml:db`
	Auth struct {
		Password string `yaml:password`
	} `yaml:auth`
}

func GetConfig() *Config {
	f, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	conf := Config{}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}

	return &conf
}
