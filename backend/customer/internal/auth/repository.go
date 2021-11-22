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
)

const (
	QUERY_GET_CUSTOMER_BY_EMAIL = "SELECT * FROM `customers` WHERE email = ?;"
	QUERY_GET_CUSTOMER_BY_PHONE = "SELECT * FROM `customers` WHERE phone_number = ?;"
	QUERY_CREATE_CUSTOMER       = "INSERT INTO `customers` (`name`, `email`, `password`, `phone_number`) VALUES (?, ?, ?, ?);"
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

func (r *repository) GetCustomerByPhone(ctx context.Context, phone string) (cust *Customer, _ error) {
	args := []interface{}{phone}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_CUSTOMER_BY_PHONE, OptsCustomerSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	cust = result.(*Customer)
	return cust, nil
}

func (r *repository) CreateCustomer(ctx context.Context, cust *Customer) (err error) {
	dbq.Tx(ctx, r.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		_, err = E(ctx, QUERY_CREATE_CUSTOMER, nil,
			cust.Name,
			cust.Email,
			cust.Password,
			cust.PhoneNumber,
		)
		if err != nil {
			return
		}
		txCommit()
	})
	return err
}
