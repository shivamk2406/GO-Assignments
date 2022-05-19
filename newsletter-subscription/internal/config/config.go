package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/shivamk2406/newsletter-subscriptions/pkg/kafka/consumer"
	"github.com/shivamk2406/newsletter-subscriptions/pkg/kafka/producer"
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
	Kafka struct {
		Brokers string `yaml:"brokers"`
	} `yaml:"kafka"`
	Producer struct {
		Servers []string `yaml:"bootstrap_servers"`
		Topic   string   `yaml:"topic"`
	} `yaml:"producer"`
	Consumer struct {
		Servers []string `yaml:"brokers"`
		GroupId string   `yaml:"groupID"`
		Topic   string   `yaml:"topic"`
	} `yaml:"consumer"`
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

func LoadProducerConfig() producer.ProducerConfig {
	var conf Config
	err := cleanenv.ReadConfig("application.yaml", &conf)
	if err != nil {
		fmt.Println(err)
		return producer.ProducerConfig{}
	}
	return producer.ProducerConfig{
		BootstrapServers: conf.Producer.Servers,
		Topic:            conf.Producer.Topic}
}

func LoadConsumerConfig() (consumer.ConsumerConfig, error) {
	var conf Config
	err := cleanenv.ReadConfig("application.yaml", &conf)
	if err != nil {
		fmt.Println(err)
		return consumer.ConsumerConfig{}, err
	}
	return consumer.ConsumerConfig{
		BootstrapServers: conf.Consumer.Servers,
		Topic:            conf.Consumer.Topic,
		Group:            conf.Consumer.GroupId}, nil
}
