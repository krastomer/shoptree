package mariadb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mariadbRepo struct {
	db *gorm.DB
}

func NewRepository() *mariadbRepo {
	dsn := "root:password@tcp(127.0.0.1:3306)/"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &mariadbRepo{db: db}
}

func (repo *mariadbRepo) GetAll() string {
	return "hello"
}
