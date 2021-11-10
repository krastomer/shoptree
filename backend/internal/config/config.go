package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig(path string) {
	viper.AddConfigPath(path)

	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {

		viper.Set("APP_PORT", ":8080")
		viper.Set("DB_HOST", "mariadb")
		viper.Set("DB_USERNAME", "root")
		viper.Set("DB_PASSWORD", "password")
		viper.Set("DB_PORT", 3306)
		viper.Set("DB_NAME", "shoptree")
		viper.Set("BCRYPT_SIZE", 8)
		viper.Set("JWT_SECRET", "september")

		viper.Set("DIRECTORY_PRODUCT", "./images/products")
		viper.Set("DIRECTORY_PAYMENT", "./images/payments")

		log.Println("Use Config Default.")
	}

}
