package order

import (
	"context"
	"database/sql"

	"github.com/rocketlaunchr/dbq/v2"
)

type repository struct {
	db *sql.DB
}

var (
	OptsProductSR        = &dbq.Options{ConcreteStruct: Product{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsOrderSR          = &dbq.Options{ConcreteStruct: Order{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsProductPendingMR = &dbq.Options{ConcreteStruct: ProductPending{}, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsAddressMR        = &dbq.Options{ConcreteStruct: Address{}, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsAddressSR        = &dbq.Options{ConcreteStruct: Address{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
)

const (
	QUERY_GET_AVAILABLE_PRODUCT              = "SELECT * FROM `products_available` WHERE id = ?;"
	QUERY_GET_ORDER_PENDING_BY_CUSTOMER_ID   = "SELECT * FROM `orders` WHERE orders.status = 'Pending' AND orders.customer_id = ?;"
	QUERY_CREATE_ORDER_PENDING               = "INSERT INTO `orders` (`customer_id`, `status`) VALUES (?, 'Pending');"
	QUERY_ADD_PRODUCT_TO_ORDER               = "INSERT INTO `products_order` (`order_id`, `product_id`) VALUES (?, ?);"
	QUERY_DELETE_PRODUCT_FROM_ORDER          = "DELETE FROM `products_order` WHERE `products_order`.`product_id` = ?;"
	QUERY_GET_PRODUCT_PENDING_BY_CUSTOMER_ID = "SELECT * FROM `products_pending` WHERE customer_id = ?"
	QUERY_GET_PRODUCT_BY_ID                  = "SELECT * FROM `products` WHERE id = ?;"
	QUERY_GET_IMAGE_PRODUCT_BY_ID            = "SELECT id FROM `images_product` WHERE product_id = ? LIMIT 1;"
	QUERY_GET_ADDRESSES_CUSTOMER             = "SELECT * FROM `addresses_customer` WHERE customer_id = ?;"
	QUERY_UPDATE_ADDRESS_ORDER               = "UPDATE `orders` SET `address_id` = ? WHERE `orders`.`id` = ?;"
	QUERY_GET_ADDRESS_CUSTOMER_BY_ID         = "SELECT * FROM `addresses_customer` WHERE id = ?;"
	QUERY_UPDATE_STATUS_ORDER                = "UPDATE `orders` SET `status` = ? WHERE `orders`.`id` = ?;"
)

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &repository{db: db}
}

func (r *repository) CreateOrderPending(ctx context.Context, custID int) (err error) {
	dbq.Tx(ctx, r.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		_, err := E(ctx, QUERY_CREATE_ORDER_PENDING, nil,
			custID,
		)
		if err != nil {
			return
		}
		txCommit()
	})
	return err
}

func (r *repository) GetOrderPendingByCustomerID(ctx context.Context, customerID int) (order *Order, _ error) {
	args := []interface{}{customerID}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_ORDER_PENDING_BY_CUSTOMER_ID, OptsOrderSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	order = result.(*Order)
	return order, nil
}

func (r *repository) AddProductToOrder(ctx context.Context, orderID int, productID int) (err error) {
	dbq.Tx(ctx, r.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		_, err = E(ctx, QUERY_ADD_PRODUCT_TO_ORDER, nil,
			orderID,
			productID,
		)
		if err != nil {
			return
		}
		txCommit()
	})

	return err
}

func (r *repository) GetAvailableProductByID(ctx context.Context, id int) (products *Product, _ error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_AVAILABLE_PRODUCT, OptsProductSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	products = result.(*Product)

	return products, nil
}

func (r *repository) DeleteProductFromOrder(ctx context.Context, productID int) (err error) {
	dbq.Tx(ctx, r.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		_, err = E(ctx, QUERY_DELETE_PRODUCT_FROM_ORDER, nil,
			productID,
		)
		if err != nil {
			return
		}
		txCommit()
	})

	return err
}

func (r *repository) GetProductPendingByCustomerID(ctx context.Context, custID int) (products []*ProductPending, _ error) {
	args := []interface{}{custID}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_PRODUCT_PENDING_BY_CUSTOMER_ID, OptsProductPendingMR, args)
	products = result.([]*ProductPending)
	if len(products) == 0 {
		return nil, sql.ErrNoRows
	}

	return products, nil
}

func (r *repository) GetProductByID(ctx context.Context, id int) (prod *Product, _ error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_PRODUCT_BY_ID, OptsProductSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	prod = result.(*Product)
	return prod, nil
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

func (r *repository) GetAddressesCustomer(ctx context.Context, custID int) (addresses []*Address, _ error) {
	args := []interface{}{custID}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_ADDRESSES_CUSTOMER, OptsAddressMR, args)
	addresses = result.([]*Address)
	if len(addresses) == 0 {
		return nil, sql.ErrNoRows
	}
	return addresses, nil
}

func (r *repository) UpdateAddressOrder(ctx context.Context, orderID, addressID int) (err error) {
	dbq.Tx(ctx, r.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		_, err = E(ctx, QUERY_UPDATE_ADDRESS_ORDER, nil,
			addressID,
			orderID,
		)
		if err != nil {
			return
		}
		txCommit()
	})

	return err
}

func (r *repository) GetAddressCustomerByID(ctx context.Context, id int) (address *Address, _ error) {
	args := []interface{}{id}

	result := dbq.MustQ(ctx, r.db, QUERY_GET_ADDRESS_CUSTOMER_BY_ID, OptsAddressSR, args)
	if result == nil {
		return nil, sql.ErrNoRows
	}
	address = result.(*Address)
	return address, nil
}

func (r *repository) UpdateStatusOrder(ctx context.Context, status string, orderID int) (err error) {
	dbq.Tx(ctx, r.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		_, err = E(ctx, QUERY_UPDATE_STATUS_ORDER, nil,
			status,
			orderID,
		)
		if err != nil {
			return
		}
		txCommit()
	})

	return err
}
