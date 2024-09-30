package config

import "github.com/spf13/viper"

func LoadConfig() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}
