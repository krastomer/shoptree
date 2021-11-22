package order

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rocketlaunchr/dbq/v2"
)

type repository struct {
	db *sql.DB
}

var (
	OptsProductSR = &dbq.Options{ConcreteStruct: Product{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
	OptsOrderSR   = &dbq.Options{ConcreteStruct: Order{}, SingleResult: true, DecoderConfig: dbq.StdTimeConversionConfig(dbq.MySQL)}
)

const (
	QUERY_GET_AVAILABLE_PRODUCT            = "SELECT * FROM `products_available` WHERE id = ?;"
	QUERY_GET_ORDER_PENDING_BY_CUSTOMER_ID = "SELECT * FROM `orders` WHERE orders.status = 'Pending' AND orders.customer_id = ?;"
	QUERY_CREATE_ORDER_PENDING             = "INSERT INTO `orders` (`customer_id`, `status`) VALUES (?, 'Pending');"
	QUERY_ADD_PRODUCT_TO_ORDER             = "INSERT INTO `products_order` (`order_id`, `product_id`) VALUES (?, ?);"
)

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &repository{db: db}
}

func (r *repository) CreateOrderPending(ctx context.Context, custID int) (err error) {
	dbq.Tx(ctx, r.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		result, err := E(ctx, QUERY_CREATE_ORDER_PENDING, nil,
			custID,
		)
		fmt.Println(result)
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
	products = result.(*Product)
	if result == nil {
		return nil, sql.ErrNoRows
	}

	return products, nil
}
