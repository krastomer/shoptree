package auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mitchellh/mapstructure"
	"github.com/rocketlaunchr/dbq/v2"
)

type mariaDBRepository struct {
	db *sql.DB
}

const (
	QUERY_GET_CUSTOMER_BY_EMAIL = "SELECT * FROM `customers` WHERE email = ?"
	QUERY_GET_EMPLOYEE_BY_EMAIL = "SELECT * FROM `employees` WHERE email = ?"
)

var (
	ErrQueryNotFound = errors.New("query not found")
)

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) GetCustomerByEmail(ctx context.Context, email string) (*Customer, error) {
	cust := &Customer{}
	args := []string{email}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_CUSTOMER_BY_EMAIL, dbq.SingleResult, args)
	if result == nil {
		return nil, ErrQueryNotFound
	}

	mapstructure.Decode(result, &cust)

	return cust, nil
}

func (r *mariaDBRepository) GetEmployeeByEmail(ctx context.Context, email string) (*Employee, error) {
	empl := &Employee{}
	args := []string{email}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_EMPLOYEE_BY_EMAIL, dbq.SingleResult, args)
	if result == nil {
		return nil, ErrQueryNotFound
	}

	mapstructure.Decode(result, &empl)

	return empl, nil
}
