package customer

import (
	"errors"

	"gorm.io/gorm"
)

type mariaDBRepository struct {
	db *gorm.DB
}

const (
	QUERY_CREATE_CUSTOMER       = "INSERT INTO `customers` (`name`, `email`, `password`, `phone_number`, `bag_level`) VALUES (?, ?, ?, ?, ?);"
	QUERY_GET_CUSTOMER_BY_EMAIL = "SELECT * FROM `customers` WHERE email = ?"
	QUERY_GET_CUSTOMER_BY_PHONE = "SELECT * FROM `customers` WHERE phone_number = ?"
)

var (
	ErrInsertFailed  = errors.New("insert failed")
	ErrQueryNotFound = errors.New("query not found")
)

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) CreateCustomer(cust *Customer) error {
	result := r.db.Exec(
		QUERY_CREATE_CUSTOMER,
		cust.Name,
		cust.Email,
		cust.Password,
		cust.PhoneNumber,
		cust.BagLevel,
	)
	if result.Error != nil {
		return ErrInsertFailed
	}
	return nil
}

func (r *mariaDBRepository) GetCustomerByEmail(email string) (*Customer, error) {
	cust := &Customer{}
	row := r.db.Raw(QUERY_GET_CUSTOMER_BY_EMAIL, email).Row()
	row.Scan(
		&cust.ID,
		&cust.Name,
		&cust.Email,
		&cust.Password,
		&cust.PhoneNumber,
		&cust.BagLevel,
		&cust.CreatedAt,
	)
	// TODO : change to compare with time
	if cust.Name == "" {
		return nil, ErrQueryNotFound
	}
	return cust, nil
}

func (r *mariaDBRepository) GetCustomerByPhone(phone string) (*Customer, error) {
	cust := &Customer{}
	row := r.db.Raw(QUERY_GET_CUSTOMER_BY_PHONE, phone).Row()
	row.Scan(
		&cust.ID,
		&cust.Name,
		&cust.Email,
		&cust.Password,
		&cust.PhoneNumber,
		&cust.BagLevel,
		&cust.CreatedAt,
	)

	// TODO : change to compare with time
	if cust.Name == "" {
		return nil, ErrQueryNotFound
	}
	return cust, nil
}
