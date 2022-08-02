package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MONGO_LOCAL_URL string `mapstructure:"MONGO_LOCAL_URL"`
	REDIS_URL       string `mapstructure:"REDIS_URL"`
	PORT            string `mapstructure:"PORT"`
}

func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
