package config

import (
	"log"

	"github.com/spf13/viper"
)

type ServerConf struct {
	Name              string
	SERVICE_USER_PORT string
}

type DatabaseConf struct {
	Driver   string
	Username string
	Password string
	Port     string
	DbName   string
	Address  string
	Schemas  string
}

type AppConfig struct {
	ServerConf
	DatabaseConf
}

func InitConfig() *AppConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../common/config")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("failed init config %v", err)
	}

	var conf *AppConfig
	err = viper.Unmarshal(&conf)

	if err != nil {
		log.Fatalf("failed unmarshall config %v", err)
	}

	log.Println(conf)

	return conf
}
