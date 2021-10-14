package mariadb

import (
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/errors"
	"github.com/krastomer/shoptree/backend/internal/models"
	"gorm.io/gorm"
)

const (
	QUERY_GET_CUSTOMER_BY_EMAIL = "SELECT * FROM `customers` WHERE email = ?"
)

func NewCustomerRepo(db *gorm.DB) entities.CustomerRepo {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) GetCustomerByEmail(email string) (*models.Customer, error) {
	cust := &models.Customer{}
	row := r.db.Raw(QUERY_GET_CUSTOMER_BY_EMAIL, email).Row()
	row.Scan(
		&cust.ID,
		&cust.Name,
		&cust.Email,
		&cust.Passwrod,
		&cust.PhoneNumber,
		&cust.CreatedAt,
	)
	if cust.Name == "" {
		return nil, errors.ErrQueryNotFound
	}
	return cust, nil
}
