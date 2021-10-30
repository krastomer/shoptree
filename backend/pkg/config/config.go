package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort    string `mapstructure:"APP_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	BcryptSize int    `mapstructure:"BCRYPT_SIZE"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`

	DirectoryProduct string `mapstructure:"DIRECTORY_PRODUCT"`
	DirectoryPayment string `mapstructure:"DIRECTORY_PAYMENT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)

	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {

		viper.Set("APP_PORT", ":8080")
		viper.Set("DB_HOST", "database")
		viper.Set("DB_USERNAME", "root")
		viper.Set("DB_PASSWORD", "password")
		viper.Set("DB_PORT", 3306)
		viper.Set("DB_NAME", "shoptree")
		viper.Set("BCRYPT_SIZE", 8)
		viper.Set("JWT_SECRET", "september")

		viper.Set("DIRECTORY_PRODUCT", "/app/image/products")
		viper.Set("DIRECTORY_PAYMENT", "/app/image/payments")

		fmt.Println("Use config Default.")

	}

	err = viper.Unmarshal(&config)
	return
}
