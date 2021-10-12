package mariadb

import (
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/errors"
	"github.com/krastomer/shoptree/backend/internal/models"
	"gorm.io/gorm"
)

const (
	QUERY_GET_CUSTOMER = "SELECT * FROM `customers` WHERE email = ?"
)

func NewCustomerRepo(db *gorm.DB) entities.CustomerRepo {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) GetCustomer(email string) (*models.Customer, error) {
	cust := &models.Customer{}
	row := r.db.Raw(QUERY_GET_CUSTOMER, email).Row()
	row.Scan(&cust.ID, &cust.Name, &cust.Email, &cust.Password, &cust.PhoneNumber, &cust.CreatedAt)
	if cust.ID == 0 || cust.Name == "" {
		return nil, errors.ErrQueryNotFound
	}
	return cust, nil
}
