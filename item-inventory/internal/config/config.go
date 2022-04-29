package config

import (
	"fmt"
	"log"
	"sync"
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
	Producer struct {
		Producer int `yaml:"producer"`
	} `yaml:"producer"`
	Consumer struct {
		Consumer int `yaml:"consumer"`
	} `yaml:"consumer"`
	Channel struct {
		BufferCapacity int `yaml:"bufferCapacity"`
	} `yaml:"channel"`
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

func LoadProducerConfig() (int, error) {
	var conf Config
	err := cleanenv.ReadConfig("application.yaml", &conf)
	if err != nil {
		fmt.Println(err)
		return conf.Producer.Producer, err
	}

	return conf.Producer.Producer, err
}

func LoadConsumerConfig() (int, error) {
	var conf Config
	err := cleanenv.ReadConfig("application.yaml", &conf)
	if err != nil {
		fmt.Println(err)
		return conf.Consumer.Consumer, err
	}

	return conf.Consumer.Consumer, err
}

func LoadChannelConfig() (int, error) {
	var conf Config
	err := cleanenv.ReadConfig("application.yaml", &conf)
	if err != nil {
		fmt.Println(err)
		return conf.Channel.BufferCapacity, err
	}

	return conf.Channel.BufferCapacity, err
}

func InitializeConfig() (Config, error) {
	var conf Config
	var confOnce sync.Once
	var err error

	confOnce.Do(func() {

		conf, err = LoadDatabaseConfig()
		if err != nil {
			log.Println(err)
		}
	})
	return conf, err
}
