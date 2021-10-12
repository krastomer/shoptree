package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectToMariaDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:password@tcp(127.0.0.1:3306)/shoptree?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
