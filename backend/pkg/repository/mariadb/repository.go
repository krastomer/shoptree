package mariadb

import (
	"github.com/krastomer/shoptree/backend/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mariadbRepo struct {
	db *gorm.DB
}

func NewRepository(dsn string) *mariadbRepo {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &mariadbRepo{db: db}
}

func (repo *mariadbRepo) GetCustomer(id int) (model.Customer, error) {
	customer := model.Customer{}
	query := "select * from customers where id = ?"
	row := repo.db.Raw(query, id).Row()
	if err := row.Err(); err != nil {
		return customer, err
	}
	row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Password, &customer.PhoneNumber, &customer.CreatedAt)
	return customer, nil
}
