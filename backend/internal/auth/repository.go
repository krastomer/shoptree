package auth

import (
	"context"
	"database/sql"

	"github.com/rocketlaunchr/dbq/v2"
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

func (r *mariaDBRepository) GetCustomerByEmail(ctx context.Context, email string) (cust *Customer, _ error) {
	args := []interface{}{email}

	opt := &dbq.Options{ConcreteStruct: Customer{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	result := dbq.MustQ(ctx, r.db, QUERY_GET_CUSTOMER_BY_EMAIL, opt, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	cust = result.(*Customer)
	return cust, nil
}

func (r *mariaDBRepository) GetEmployeeByEmail(ctx context.Context, email string) (empl *Employee, _ error) {
	args := []interface{}{email}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_CUSTOMER_BY_EMAIL, dbq.SingleResult, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	empl = result.(*Employee)
	return empl, nil
}
