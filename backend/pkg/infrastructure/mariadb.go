package infrastructure

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB_URL      = "0.0.0.0"
	DB_USER     = "root"
	DB_PASSWORD = "password"
	DB_PORT     = "3306"
	DB_DEFAULT  = "shoptree"
)

func init() {
	if viper.GetBool("FROM_COMPOSE") {
		DB_URL = "database"
	}
}

func connectToMariaDB() (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER,
		DB_PASSWORD,
		DB_URL,
		DB_PORT,
		DB_DEFAULT,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(3 * time.Minute)

	return db, nil
}
