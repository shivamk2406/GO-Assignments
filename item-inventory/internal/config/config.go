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
		RoutineCount int `yaml:"routineCount"`
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

func LoadRoutineConfig() (int, error) {
	var conf Config
	err := cleanenv.ReadConfig("application.yaml", &conf)
	if err != nil {
		fmt.Println(err)
		return conf.Routine.RoutineCount, err
	}
	return conf.Routine.RoutineCount, err
}
