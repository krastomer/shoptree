package infrastructure

import (
	"fmt"
	"time"

	"github.com/krastomer/shoptree/backend/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectToMariaDB(config config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
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
