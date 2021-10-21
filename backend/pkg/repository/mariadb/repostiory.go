package mariadb

import "gorm.io/gorm"

type mariaDBRepository struct {
	db *gorm.DB
}
