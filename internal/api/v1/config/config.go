package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Dev                 Environment `yaml:"dev"`
	Prod                Environment `yaml:"prod"`
	SelectedEnvironment string      `yaml:"selected_environment"`
}

type Environment struct {
	Database DatabaseConfig `yaml:"database"`
	AuthFile string         `yaml:"authfile"`
	Port string       `yaml:"port"`
}

type DatabaseConfig struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func ReadConfig(cofigPath string) (*Environment, error) {
	var c Config

	yamlFile, err := ioutil.ReadFile(cofigPath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatal(err)
	}

	if c.SelectedEnvironment == "dev" {
		return &c.Dev, nil
	}

	return &c.Prod, nil
}
