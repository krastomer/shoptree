package mariadb

import (
	"github.com/krastomer/shoptree/backend/pkg/auth"
	"gorm.io/gorm"
)

const (
	QUERY_GET_CUSTOMER_BY_EMAIL = "SELECT * FROM `customers` WHERE email = ?"
	QUERY_GET_CUSTOMER_BY_PHONE = "SELECT * FROM `customers` WHERE phone_number = ?"
	QUERY_GET_EMPLOYEE_BY_EMAIL = "SELECT * FROM `employees` WHERE email = ?"
	QUERY_REGISTER_CUSTOMER     = "INSERT INTO `customers` (`name`, `email`, `password`, `phone_number`, `bag_level`) VALUES (?, ?, ?, ?, ?);"
)

func NewAuthRepository(db *gorm.DB) auth.AuthRepository {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) GetCustomerByEmail(email string) (*auth.Customer, error) {
	cust := &auth.Customer{}
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
	if cust.Name == "" {
		return nil, ErrQueryNotFound
	}
	return cust, nil
}

func (r *mariaDBRepository) GetCustomerByPhone(phone string) (*auth.Customer, error) {
	cust := &auth.Customer{}
	row := r.db.Raw(QUERY_GET_CUSTOMER_BY_PHONE, phone).Row()
	row.Scan(
		&cust.ID,
		&cust.Name,
		&cust.Email,
		&cust.Password,
		&cust.PhoneNumber,
		&cust.CreatedAt,
	)
	if cust.Name == "" {
		return nil, ErrQueryNotFound
	}
	return cust, nil
}

func (r *mariaDBRepository) GetEmployeeByEmail(email string) (*auth.Employee, error) {
	empl := &auth.Employee{}
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
	if empl.Name == "" {
		return nil, ErrQueryNotFound
	}
	return empl, nil
}

func (r *mariaDBRepository) RegisterCustomer(cust *auth.Customer) error {
	result := r.db.Exec(
		QUERY_REGISTER_CUSTOMER,
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
