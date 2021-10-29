package infrastructure

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB_URL      = "192.168.1.34"
	DB_USER     = "root"
	DB_PASSWORD = "password"
	DB_PORT     = "3306"
	DB_DEFAULT  = "shoptree"
)

func connectToMariaDB() (*gorm.DB, error) {

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	DB_USER,
	// 	DB_PASSWORD,
	// 	DB_URL,
	// 	DB_PORT,
	// 	DB_DEFAULT,
	// )

	dsn := "root:password@tcp(database:3306)/shoptree?charset=utf8mb4&parseTime=True&loc=Local"
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
