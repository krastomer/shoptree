package mariadb

import (
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/models"
	"gorm.io/gorm"
)

func NewCustomerRepo(db *gorm.DB) entities.CustomerRepo {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) GetCustomer(email string) (*models.Customer, error) {
	cust := &models.Customer{}
	query := "SELECT * FROM `customers` WHERE email = ?"
	row := r.db.Raw(query, email).Row()
	row.Scan(&cust.ID, &cust.Name, &cust.Email, &cust.Password, &cust.PhoneNumber, &cust.CreatedAt)
	return cust, nil
}
