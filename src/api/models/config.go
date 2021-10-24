package models

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Endpoint string `yaml:"endpoint"`
		Key string `yaml:"key"`
	} `yaml:"database"`
}