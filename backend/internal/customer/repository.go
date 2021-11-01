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
	QUERY_GET_CUSTOMER_BY_ID    = "SELECT * FROM `customers` WHERE id = ?"
	QUERY_GET_ADDRESSES         = "SELECT * FROM `address_customers` WHERE customer_id = ?"
)

var (
	ErrInsertFailed        = errors.New("insert failed")
	ErrQueryNotFound       = errors.New("query not found")
	ErrInternalServerError = errors.New("internal server error")
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

func (r *mariaDBRepository) GetAddresses(id int) ([]*Address, error) {
	var addresses []*Address
	rows, err := r.db.Raw(QUERY_GET_ADDRESSES, id).Rows()
	if err != nil {
		return nil, ErrInternalServerError
	}
	defer rows.Close()
	for rows.Next() {
		address := &Address{}
		r.db.ScanRows(rows, address)
		addresses = append(addresses, address)
	}
	return addresses, nil
}

func (r *mariaDBRepository) GetCustomerByID(id int) (*Customer, error) {
	cust := &Customer{}
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

	// TODO : change to compare with time
	if cust.Name == "" {
		return nil, ErrQueryNotFound
	}
	return cust, nil
}
