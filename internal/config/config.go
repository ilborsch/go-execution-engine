package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env     string        `yaml:"env" env-default:"local"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	path := getConfigPath()
	return MustLoadByPath(path)
}

func getConfigPath() string {
	var path string
	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("GO_EXECUTOR_CONFIG_PATH")
	}
	return path
}

func MustLoadByPath(path string) *Config {
	if path == "" {
		panic("config path is empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file with given path does not exist: " + path)
	}
	var config Config

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		panic("error parsing config file with given path: " + path + " " + err.Error())
	}
	return &config
}
