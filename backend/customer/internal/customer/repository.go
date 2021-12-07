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
	OptsAddressSR  = &dbq.Options{ConcreteStruct: Address{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsCustomerSR = &dbq.Options{ConcreteStruct: Customer{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsOrderMR    = &dbq.Options{ConcreteStruct: Order{}, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsProductMR  = &dbq.Options{ConcreteStruct: Product{}, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
)

const (
	QUERY_GET_ADDRESSES_CUSTOMER     = "SELECT * FROM `addresses_customer` WHERE customer_id = ?;"
	QUERY_GET_CUSTOMER_BY_ID         = "SELECT * FROM `customers` WHERE id = ?;"
	QUERY_CREATE_ADDRESS_CUSTOMER    = "INSERT INTO `addresses_customer` (`customer_id`, `name`, `phone_number`, `address_line`, `country`, `state`, `city`, `district`, `postal_code`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);"
	QUERY_GET_ADDRESS_CUSTOMER_BY_ID = "SELECT * FROM `addresses_customer` WHERE id = ?;"
	QUERY_DELETE_ADDRESS_CUSTOMER    = "DELETE FROM `addresses_customer` WHERE id = ?"
	QUERY_GET_ORDERS_CUSTOMER        = "SELECT * FROM `orders` WHERE `orders`.`status` NOT IN ('Undefined','Pending','Failed') AND orders.customer_id = ?;"
	QUERY_GET_PRODUCTS_BY_ORDER_ID   = "SELECT * FROM `products` WHERE id IN (SELECT products_order.product_id FROM products_order JOIN `orders` ON products_order.order_id = orders.id WHERE orders.id = ?);"
	QUERY_GET_IMAGE_PRODUCT_BY_ID    = "SELECT id FROM `images_product` WHERE product_id = ? LIMIT 1;"
	QUERY_GET_ADDRESS_BY_ORDER_ID    = "SELECT * FROM addresses_customer WHERE addresses_customer.id = (SELECT orders.address_id FROM orders WHERE orders.id = ?);"
	QUERY_GET_PAYMENT_BY_ORDER_ID    = "SELECT image_path FROM payments WHERE order_id = ?;"
)

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &repository{db: db}
}

func (r *repository) GetAddressCustomerByID(ctx context.Context, id int) (address *Address, err error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_ADDRESS_CUSTOMER_BY_ID, OptsAddressSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	address = result.(*Address)
	return address, nil
}

func (r *repository) DeleteAddressCustomer(ctx context.Context, id int) (err error) {
	dbq.Tx(ctx, r.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		_, err = E(ctx, QUERY_DELETE_ADDRESS_CUSTOMER, nil, id)
		if err != nil {
			return
		}
		txCommit()
	})
	return err
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

func (r *repository) CreateAddressCustomer(ctx context.Context, address *Address) (err error) {
	dbq.Tx(ctx, r.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		_, err = E(ctx, QUERY_CREATE_ADDRESS_CUSTOMER, nil,
			address.CustomerID,
			address.Name,
			address.PhoneNumber,
			address.AddressLine,
			address.Country,
			address.State,
			address.City,
			address.District,
			address.PostalCode,
		)
		if err != nil {
			return
		}
		txCommit()
	})
	return err
}

func (r *repository) GetOrdersCustomer(ctx context.Context, custID int) (orders []*Order, err error) {
	args := []interface{}{custID}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_ORDERS_CUSTOMER, OptsOrderMR, args)
	orders = result.([]*Order)
	if len(orders) == 0 {
		return nil, sql.ErrNoRows
	}
	return orders, nil
}

func (r *repository) GetProductsByOrderID(ctx context.Context, orderID int) (products []*Product, err error) {
	args := []interface{}{orderID}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_PRODUCTS_BY_ORDER_ID, OptsProductMR, args)
	products = result.([]*Product)
	if len(products) == 0 {
		return nil, sql.ErrNoRows
	}
	return products, nil
}

func (r *repository) GetImageProductByID(ctx context.Context, id int) (path int, _ error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_IMAGE_PRODUCT_BY_ID, dbq.SingleResult, args)
	if result == nil {
		return -1, sql.ErrNoRows
	}
	path = int(result.(map[string]interface{})["id"].(int32))
	return path, nil
}

func (r *repository) GetAddressByOrderID(ctx context.Context, orderID int) (address *Address, err error) {
	args := []interface{}{orderID}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_ADDRESS_BY_ORDER_ID, OptsAddressSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	address = result.(*Address)
	return address, nil
}

func (r *repository) GetPaymentByOrderID(ctx context.Context, orderID int) (path string, err error) {
	args := []interface{}{orderID}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_PAYMENT_BY_ORDER_ID, dbq.SingleResult, args)
	if result == nil {
		return "", sql.ErrNoRows
	}
	path = result.(map[string]interface{})["image_path"].(string)

	return path, nil
}
