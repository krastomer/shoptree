package customer

import (
	"context"
	"database/sql"

	"github.com/rocketlaunchr/dbq/v2"
)

type mariaDBRepository struct {
	db *sql.DB
}

var (
	OptsCustomerSR        = &dbq.Options{ConcreteStruct: Customer{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsAddressCustomerMR = &dbq.Options{ConcreteStruct: Address{}, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
)

const (
	QUERY_CREATE_CUSTOMER        = "INSERT INTO `customers` (`name`, `email`, `password`, `phone_number`) VALUES (?, ?, ?, ?);"
	QUERY_GET_CUSTOMER_BY_EMAIL  = "SELECT * FROM `customers` WHERE email = ?"
	QUERY_GET_CUSTOMER_BY_ID     = "SELECT * FROM `customers` WHERE id = ?"
	QUERY_GET_CUSTOMER_BY_PHONE  = "SELECT * FROM `customers` WHERE phone_number = ?"
	QUERY_GET_ADDRESSES_CUSTOMER = "SELECT * FROM `addresses_customer` WHERE customer_id = ?"

	QUERY_CREATE_ADDRESS = "INSERT INTO `address_customers` ( `customer_id`, `name`, `phone_number`, `address_line`, `country`, `state`, `city`, `district`, `postal_code`) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?);"

	QUERY_GET_INVOICES = "SELECT * FROM `invoices` WHERE customer_id = ?;"
)

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &mariaDBRepository{db: db}
}

func (r *mariaDBRepository) CreateCustomer(ctx context.Context, cust *CustomerRequest) (err error) {
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
	return nil
}

func (r *mariaDBRepository) GetCustomerByEmail(ctx context.Context, email string) (cust *Customer, _ error) {
	args := []interface{}{email}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_CUSTOMER_BY_EMAIL, OptsCustomerSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	cust = result.(*Customer)
	return cust, nil
}

func (r *mariaDBRepository) GetCustomerByID(ctx context.Context, id int) (cust *Customer, _ error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_CUSTOMER_BY_ID, OptsCustomerSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	cust = result.(*Customer)
	return cust, nil
}

func (r *mariaDBRepository) GetCustomerByPhone(ctx context.Context, phone string) (cust *Customer, _ error) {
	args := []interface{}{phone}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_CUSTOMER_BY_PHONE, OptsCustomerSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	cust = result.(*Customer)
	return cust, nil
}

func (r *mariaDBRepository) GetAddressesCustomer(ctx context.Context, id int) (addresses []*Address, err error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_ADDRESSES_CUSTOMER, OptsAddressCustomerMR, args)
	addresses = result.([]*Address)
	// nil didn't work with array
	if len(addresses) == 0 {
		return nil, sql.ErrNoRows
	}

	return addresses, nil
}
