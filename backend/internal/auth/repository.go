package auth

import (
	"context"
	"database/sql"
)

type mariaDBRepository struct {
	db *sql.DB
}

const (
	QUERY_GET_CUSTOMER_BY_EMAIL = "SELECT * FROM `customers` WHERE email = ?"
	QUERY_GET_EMPLOYEE_BY_EMAIL = "SELECT * FROM `employees` WHERE email = ?"
)

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) GetCustomerByEmail(ctx context.Context, email string) (*Customer, error) {
	cust := &Customer{}

	stmt, err := r.db.PrepareContext(ctx, QUERY_GET_CUSTOMER_BY_EMAIL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, email).Scan(
		&cust.ID,
		&cust.Name,
		&cust.Email,
		&cust.Password,
		&cust.PhoneNumber,
		&cust.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return cust, nil
}

func (r *mariaDBRepository) GetEmployeeByEmail(ctx context.Context, email string) (*Employee, error) {
	empl := &Employee{}

	stmt, err := r.db.PrepareContext(ctx, QUERY_GET_EMPLOYEE_BY_EMAIL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, email).Scan(
		&empl.ID,
		&empl.Name,
		&empl.Email,
		&empl.Password,
		&empl.PhoneNumber,
		&empl.Level,
		&empl.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return empl, nil
}
