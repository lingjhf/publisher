package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseName string `mapstructure:"database_name"`
	Host         string `mapstructure:"host"`
	Port         uint   `mapstructure:"port"`
}

func Read() *Config {
	v := viper.New()
	v.SetConfigName("publisher.config")
	v.AddConfigPath(".")
	v.SetConfigType("json")
	c := &Config{
		DatabaseName: "publish",
		Host:         "localhost",
		Port:         3000,
	}
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			v.WatchConfig()
		}
	} else {
		v.Unmarshal(c)
	}
	return c
}
