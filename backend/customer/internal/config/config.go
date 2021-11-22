package config

import "github.com/spf13/viper"

func LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)

	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	return err
}
