package auth

import (
	"errors"

	"gorm.io/gorm"
)

type mariaDBRepository struct {
	db *gorm.DB
}

const (
	QUERY_GET_CUSTOMER_BY_EMAIL = "SELECT * FROM `customers` WHERE email = ?"
	QUERY_GET_EMPLOYEE_BY_EMAIL = "SELECT * FROM `employees` WHERE email = ?"
)

var (
	ErrQueryNotFound = errors.New("query not found")
)

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &mariaDBRepository{db: db}
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

func (r *mariaDBRepository) GetEmployeeByEmail(email string) (*Employee, error) {
	empl := &Employee{}
	row := r.db.Raw(QUERY_GET_EMPLOYEE_BY_EMAIL, email).Row()
	row.Scan(
		&empl.ID,
		&empl.Name,
		&empl.Email,
		&empl.Password,
		&empl.PhoneNumber,
		&empl.Level,
		&empl.CreatedAt,
	)
	// TODO : change to compare with time
	if empl.Name == "" {
		return nil, ErrQueryNotFound
	}
	return empl, nil
}
