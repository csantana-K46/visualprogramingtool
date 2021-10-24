package config


import (
	m "api/models"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type appConfig struct {
	Configurations m.Config
}

func NewAppConfig() appConfig {
	f, err := os.Open("env.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	var cfg m.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		processError(err)
	}
	c := appConfig {cfg}
	return c
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}