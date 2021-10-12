package mariadb

import (
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/errors"
	"github.com/krastomer/shoptree/backend/internal/models"
	"gorm.io/gorm"
)

const (
	QUERY_GET_CUSTOMER_BY_EMAIL = "SELECT * FROM `customers` WHERE email = ?"
	QUERY_GET_CUSTOMER_BY_PHONE = "SELECT * FROM `customers` WHERE phone_number = ?"
	QUERY_CREATE_CUSTOMER       = "INSERT INTO `customers` (`name`, `email`, `password`, `phone_number`) VALUES (?, ?, ?, ?);"
)

func NewCustomerRepo(db *gorm.DB) entities.CustomerRepo {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) GetCustomerByEmail(email string) (*models.Customer, error) {
	cust := &models.Customer{}
	row := r.db.Raw(QUERY_GET_CUSTOMER_BY_EMAIL, email).Row()
	row.Scan(&cust.ID, &cust.Name, &cust.Email, &cust.Password, &cust.PhoneNumber, &cust.CreatedAt)
	if cust.ID == 0 || cust.Name == "" {
		return nil, errors.ErrQueryNotFound
	}
	return cust, nil
}

func (r *mariaDBRepository) GetCustomerByPhone(phone string) (*models.Customer, error) {
	cust := &models.Customer{}
	row := r.db.Raw(QUERY_GET_CUSTOMER_BY_PHONE, phone).Row()
	row.Scan(&cust.ID, &cust.Name, &cust.Email, &cust.Password, &cust.PhoneNumber, &cust.CreatedAt)
	if cust.ID == 0 || cust.Name == "" {
		return nil, errors.ErrQueryNotFound
	}
	return cust, nil
}

func (r *mariaDBRepository) CreateCustomer(cust *models.CustomerProfile) error {
	row := r.db.Exec(QUERY_CREATE_CUSTOMER, cust.Name, cust.Email, cust.Password, cust.PhoneNumber)
	if row.Error != nil {
		return errors.ErrInsertFailed
	}
	return nil
}

func (r *mariaDBRepository) UpdateCustomerPassword(new string) error {
	return nil
}
