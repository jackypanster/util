package util

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
}

func NewConfig() *Config {
	viper.SetConfigName("cfg")
	viper.AddConfigPath("/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("error %s", err.Error())
	}
	return &Config{}
}

func (self *Config) GetString(key string) string {
	return viper.GetString(key)
}

func (self *Config) GetInt(key string) int {
	return viper.GetInt(key)
}

func (self *Config) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (self *Config) GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}
