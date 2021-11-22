package customer

import (
	"context"
	"database/sql"

	"github.com/rocketlaunchr/dbq/v2"
)

type repository struct {
	db *sql.DB
}

var (
	OptsAddressMR  = &dbq.Options{ConcreteStruct: Address{}, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsCustomerSR = &dbq.Options{ConcreteStruct: Customer{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
)

const (
	QUERY_GET_ADDRESSES_CUSTOMER = "SELECT * FROM `addresses_customer` WHERE customer_id = ?;"
	QUERY_GET_CUSTOMER_BY_ID     = "SELECT * FROM `customers` WHERE id = ?;"
)

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &repository{db: db}
}

func (r *repository) GetAddressesCustomer(ctx context.Context, custID int) (addresses []*Address, _ error) {
	args := []interface{}{custID}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_ADDRESSES_CUSTOMER, OptsAddressMR, args)
	addresses = result.([]*Address)
	if len(addresses) == 0 {
		return nil, sql.ErrNoRows
	}
	return addresses, nil
}

func (r *repository) GetCustomerByID(ctx context.Context, id int) (cust *Customer, _ error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_CUSTOMER_BY_ID, OptsCustomerSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	cust = result.(*Customer)
	return cust, nil
}
