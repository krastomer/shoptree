package infrastructure

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectToMariaDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:password@tcp(database:3306)/shoptree?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(3 * time.Minute)

	return db, nil
}
