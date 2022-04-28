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
	Routine struct {
		Consumer int `yaml:"consumer"`
		Producer int `yaml:"producer"`
	} `yaml:"routine"`
}

type Routine struct {
	Routine struct {
		Consumer int `yaml:"consumer"`
		Producer int `yaml:"producer"`
	} `yaml:"routine"`
}

func LoadDatabaseConfig() (Config, error) {
	var conf Config
	err := cleanenv.ReadConfig("application.yaml", &conf)
	if err != nil {
		fmt.Println(err)
		return conf, err
	}
	return conf, nil
}

func LoadRoutineConfig() (int, int, error) {
	var conf Config
	err := cleanenv.ReadConfig("application.yaml", &conf)
	if err != nil {
		fmt.Println(err)
		return conf.Routine.Consumer, conf.Routine.Producer, err
	}

	return conf.Routine.Consumer, conf.Routine.Producer, err
}
