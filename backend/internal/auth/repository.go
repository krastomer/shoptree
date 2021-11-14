package auth

import (
	"context"
	"database/sql"

	"github.com/rocketlaunchr/dbq/v2"
)

type repository struct {
	db *sql.DB
}

var (
	OptsCustomerSR = &dbq.Options{ConcreteStruct: Customer{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsEmployeeSR = &dbq.Options{ConcreteStruct: Employee{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
)

const (
	QUERY_GET_CUSTOMER_BY_EMAIL = "SELECT * FROM `customers` WHERE email = ?"
	QUERY_GET_EMPLOYEE_BY_EMAIL = "SELECT * FROM `employees` WHERE email = ?"
)

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &repository{db: db}
}

func (r *repository) GetCustomerByEmail(ctx context.Context, email string) (cust *Customer, _ error) {
	args := []interface{}{email}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_CUSTOMER_BY_EMAIL, OptsCustomerSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	cust = result.(*Customer)
	return cust, nil
}

func (r *repository) GetEmployeeByEmail(ctx context.Context, email string) (empl *Employee, _ error) {
	args := []interface{}{email}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_EMPLOYEE_BY_EMAIL, OptsEmployeeSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	empl = result.(*Employee)
	return empl, nil
}
