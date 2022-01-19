package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const configPath = "config.yml"

type Cfg struct {
	UserName     string `yaml:"user_name"`
	Password     string `yaml:"password"`
	Addr         string `yaml:"addr"`
	Port         int    `yaml:"port"`
	Database     string `yaml:"database"`
	MaxLifetime  int    `yaml:"max_lifetime"`
	MaxOpenConns int    `yaml:"max_open_conns"`
}

var AppConfig Cfg

func ReadConfig() {
	f, err := os.Open(configPath)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&AppConfig)

	if err != nil {
		log.Println(err)
	}
}
