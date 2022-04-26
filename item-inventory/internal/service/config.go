package service

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/shivamk2406/item-inventory/database"
)

func LoadAppConfig() (error, database.Config) {

	var conf database.Config
	err := cleanenv.ReadConfig("application.yaml", &conf)
	if err != nil {
		fmt.Println(err)
		return err, conf
	}

	return nil, conf
}
