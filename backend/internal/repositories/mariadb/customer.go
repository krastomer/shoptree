package mariadb

import (
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/models"
	"gorm.io/gorm"
)

func NewCustomerRepo(db *gorm.DB) entities.CustomerRepo {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) GetCustomer(id int) (*models.Customer, error) {
	cust := &models.Customer{}
	query := "SELECT name FROM `customers` WHERE id = ?"
	r.db.Raw(query, id).Scan(&cust.Name)
	return cust, nil
}
