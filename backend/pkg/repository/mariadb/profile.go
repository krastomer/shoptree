package mariadb

import (
	"github.com/krastomer/shoptree/backend/pkg/profile"
	"gorm.io/gorm"
)

const (
	QUERY_GET_CUSTOMER_BY_ID = "SELECT * FROM `customers` WHERE id = ?"
	QUERY_GET_ADDRESSES      = "SELECT * FROM `address_customers` WHERE customer_id = ?"
)

func NewProfileRepository(db *gorm.DB) profile.ProfileRepository {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) GetCustomerByID(id uint32) (*profile.Customer, error) {
	cust := &profile.Customer{}
	row := r.db.Raw(QUERY_GET_CUSTOMER_BY_ID, id).Row()
	row.Scan(
		&cust.ID,
		&cust.Name,
		&cust.Email,
		&cust.Password,
		&cust.PhoneNumber,
		&cust.BagLevel,
		&cust.CreatedAt,
	)
	if cust.Name == "" {
		return nil, ErrQueryNotFound
	}
	return cust, nil
}

func (r *mariaDBRepository) GetAddresses(id uint32) ([]*profile.Address, error) {
	var addresses []*profile.Address
	rows, err := r.db.Raw(QUERY_GET_ADDRESSES, id).Rows()
	if err != nil {
		return nil, ErrInternalServerError
	}
	defer rows.Close()
	for rows.Next() {
		address := &profile.Address{}
		r.db.ScanRows(rows, address)
		addresses = append(addresses, address)
	}
	return addresses, nil
}
