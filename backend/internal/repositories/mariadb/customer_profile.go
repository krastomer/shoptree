package mariadb

import (
	"github.com/krastomer/shoptree/backend/internal/entities"
	"github.com/krastomer/shoptree/backend/internal/errors"
	"github.com/krastomer/shoptree/backend/internal/models"
	"gorm.io/gorm"
)

const (
	QUERY_GET_CUSTOMER_BY_ID = "SELECT * FROM `customers` WHERE id = ?"
	QUERY_GET_ADDRESSES      = "SELECT * FROM `address_customers` WHERE customer_id = ?"
)

func NewCustomerProfileRepo(db *gorm.DB) entities.CustomerProfileRepo {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) GetCustomerByID(id uint32) (*models.Customer, error) {
	cust := &models.Customer{}
	row := r.db.Raw(QUERY_GET_CUSTOMER_BY_ID, id).Row()
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

func (r *mariaDBRepository) GetAddresses(cust_id uint32) ([]*models.Address, error) {
	var addresses []*models.Address
	rows, err := r.db.Raw(QUERY_GET_ADDRESSES, cust_id).Rows()
	if err != nil {
		return nil, errors.ErrQueryNotFound
	}
	defer rows.Close()
	for rows.Next() {
		address := &models.Address{}
		r.db.ScanRows(rows, address)
		addresses = append(addresses, address)
	}
	return addresses, nil
}
