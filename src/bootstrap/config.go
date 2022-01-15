package bootstrap

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server Server `yaml:"server"`
	Bot    Bot    `yaml:"bot"`
	Router Router `yaml:"router"`
	DB     DB     `yaml:"database"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Bot struct {
	Token string `yaml:"token"`
}

type Router struct {
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	WlURL    string `yaml:"wl_url"`
}

type DB struct {
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func New() (*Config, error) {
	config := new(Config)

	f, err := os.ReadFile("./config.yml")
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(f, config); err != nil {
		return nil, err
	}

	return config, nil
}
