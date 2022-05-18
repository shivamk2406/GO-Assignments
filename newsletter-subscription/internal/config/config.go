package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Database struct {
		User                  string        `yaml:"user"`
		Password              string        `yaml:"password"`
		Host                  string        `yaml:"host"`
		Name                  string        `yaml:"name"`
		MaxIdleConnections    int           `yaml:"maxIdleConnections"`
		MaxOpenConnections    int           `yaml:"maxOpenConnections"`
		MaxConnectionLifeTime time.Duration `yaml:"maxConnectionLifeTime"`
		MaxConnectionIdleTime time.Duration `yaml:"maxConnectionIdleTime"`
		DisableTLS            bool          `yaml:"disableTLS"`
	} `yaml:"database"`
	Grpc struct {
		Port    string `yaml:"port"`
		Network string `yaml:"network"`
	} `yaml:"grpc"`
}

func LoadDatabaseConfig() (Config, error) {
	var conf Config
	err := cleanenv.ReadConfig("application.yaml", &conf)
	if err != nil {
		fmt.Println(err)
		return conf, err
	}
	fmt.Println(conf)
	return conf, nil
}

func LoadGrpcConfig() (string, string, error) {
	var conf Config
	err := cleanenv.ReadConfig("application.yaml", &conf)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	return conf.Grpc.Port, conf.Grpc.Network, nil
}
